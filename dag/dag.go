// Package dag 提供一个有向无环图（DAG）任务编排工具。
//
// 通过声明任务及其依赖关系构建 DAG，执行时自动进行拓扑排序、环检测，
// 并在满足依赖约束的前提下最大化并行执行。支持 context 取消、并发上限、
// 出错快速失败或继续执行，以及任务间的数据传递。
package dag

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	gerrors "github.com/Ccmuyu/gopkg/errors"
)

// Results 保存各任务的输出，键为任务名。
type Results map[string]any

// Task 任务执行函数。
// deps 为该任务【直接依赖】任务的输出集合（键为依赖任务名）；
// 返回值作为该任务的输出，供下游任务通过 deps 读取。
// 实现应尊重 ctx 取消，以便在快速失败时及时退出。
type Task func(ctx context.Context, deps Results) (any, error)

// node 内部节点
type node struct {
	name string
	deps []string
	fn   Task
}

// DAG 任务编排图。构建期（Add）非并发安全；构建完成后 Run 可并发执行内部任务。
type DAG struct {
	nodes map[string]*node
	order []string // 记录插入顺序，保证拓扑输出稳定
	err   error    // 构建期错误，延迟到 Validate/Run 暴露
}

// New 创建空的 DAG。
func New() *DAG {
	return &DAG{nodes: make(map[string]*node)}
}

// Add 添加一个任务：name 为唯一任务名，deps 为其依赖的任务名列表，fn 为执行函数。
// 返回 *DAG 以支持链式调用；构建期的错误（重名、空名、nil 函数）会被记录，
// 并在 Validate / Run / TopoOrder 时返回。
func (d *DAG) Add(name string, deps []string, fn Task) *DAG {
	if d.err != nil {
		return d
	}
	if name == "" {
		d.err = gerrors.New("dag: task name is empty")
		return d
	}
	if fn == nil {
		d.err = fmt.Errorf("dag: task %q has nil func", name)
		return d
	}
	if _, ok := d.nodes[name]; ok {
		d.err = fmt.Errorf("dag: duplicate task %q", name)
		return d
	}
	d.nodes[name] = &node{
		name: name,
		deps: append([]string(nil), deps...),
		fn:   fn,
	}
	d.order = append(d.order, name)
	return d
}

// Len 返回任务数量。
func (d *DAG) Len() int { return len(d.nodes) }

// TaskNames 按插入顺序返回所有任务名。
func (d *DAG) TaskNames() []string {
	return append([]string(nil), d.order...)
}

// Validate 校验图的合法性：构建期错误、依赖缺失、以及是否存在环。
func (d *DAG) Validate() error {
	_, err := d.TopoOrder()
	return err
}

// TopoOrder 返回一个拓扑排序序列（依赖在前）。
// 若存在构建错误、依赖缺失或环，返回相应错误。结果对相同输入是稳定的。
func (d *DAG) TopoOrder() ([]string, error) {
	if d.err != nil {
		return nil, d.err
	}

	indeg := make(map[string]int, len(d.nodes))
	for _, name := range d.order {
		indeg[name] = 0
	}
	adj := make(map[string][]string, len(d.nodes))
	for _, name := range d.order {
		for _, dep := range d.nodes[name].deps {
			if _, ok := d.nodes[dep]; !ok {
				return nil, fmt.Errorf("dag: task %q depends on unknown task %q", name, dep)
			}
			adj[dep] = append(adj[dep], name)
			indeg[name]++
		}
	}

	queue := make([]string, 0, len(d.nodes))
	for _, name := range d.order {
		if indeg[name] == 0 {
			queue = append(queue, name)
		}
	}

	out := make([]string, 0, len(d.nodes))
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		out = append(out, n)
		for _, m := range adj[n] {
			indeg[m]--
			if indeg[m] == 0 {
				queue = append(queue, m)
			}
		}
	}

	if len(out) != len(d.nodes) {
		return nil, fmt.Errorf("dag: cycle detected among %s", strings.Join(remaining(indeg), ", "))
	}
	return out, nil
}

// remaining 返回仍有未满足依赖（indeg>0）的任务名，用于环报错提示。
func remaining(indeg map[string]int) []string {
	var names []string
	for name, deg := range indeg {
		if deg > 0 {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return names
}

// Option 执行选项
type Option func(*runConfig)

type runConfig struct {
	maxConcurrency  int
	continueOnError bool
}

// WithMaxConcurrency 限制同时运行的任务数量；n <= 0 表示不限制。
func WithMaxConcurrency(n int) Option {
	return func(c *runConfig) { c.maxConcurrency = n }
}

// WithContinueOnError 某任务失败时不中断整个图：
// 仅跳过其（直接/间接）下游任务，其余互不依赖的任务继续执行。
// 默认行为为快速失败（首个错误即取消其余未开始的任务）。
func WithContinueOnError() Option {
	return func(c *runConfig) { c.continueOnError = true }
}

// Run 执行整个 DAG，返回所有成功任务的输出集合与聚合错误。
//
// 依赖满足即可并行执行；默认快速失败：任一任务出错会取消 ctx，
// 尚未开始的任务将被跳过。使用 WithContinueOnError 可改为尽力执行。
// 无论成功与否都会返回已完成任务的部分结果。
func (d *DAG) Run(ctx context.Context, opts ...Option) (Results, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}

	cfg := &runConfig{}
	for _, o := range opts {
		o(cfg)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var sem chan struct{}
	if cfg.maxConcurrency > 0 {
		sem = make(chan struct{}, cfg.maxConcurrency)
	}

	results := make(Results, len(d.nodes))
	failed := make(map[string]struct{}, len(d.nodes))
	merr := &gerrors.MultiError{}
	var mu sync.Mutex

	done := make(map[string]chan struct{}, len(d.nodes))
	for _, name := range d.order {
		done[name] = make(chan struct{})
	}

	markFailed := func(name string) {
		mu.Lock()
		failed[name] = struct{}{}
		mu.Unlock()
	}

	var wg sync.WaitGroup
	for _, name := range d.order {
		name := name
		t := d.nodes[name]
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(done[name])

			// 等待所有依赖完成（或 ctx 取消）
			for _, dep := range t.deps {
				select {
				case <-done[dep]:
				case <-ctx.Done():
					markFailed(name)
					return
				}
			}

			// 汇总依赖输出，并判断是否需要跳过
			mu.Lock()
			skip := ctx.Err() != nil
			depResults := make(Results, len(t.deps))
			for _, dep := range t.deps {
				if _, bad := failed[dep]; bad {
					skip = true
					break
				}
				depResults[dep] = results[dep]
			}
			mu.Unlock()
			if skip {
				markFailed(name)
				return
			}

			// 并发上限
			if sem != nil {
				select {
				case sem <- struct{}{}:
					defer func() { <-sem }()
				case <-ctx.Done():
					markFailed(name)
					return
				}
			}

			out, err := t.fn(ctx, depResults)

			mu.Lock()
			if err != nil {
				failed[name] = struct{}{}
				merr.Append(fmt.Errorf("dag: task %q: %w", name, err))
				mu.Unlock()
				if !cfg.continueOnError {
					cancel()
				}
				return
			}
			results[name] = out
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results, merr.ErrorOrNil()
}

// DOT 输出 Graphviz DOT 格式，便于可视化依赖关系。
func (d *DAG) DOT() string {
	var b strings.Builder
	b.WriteString("digraph dag {\n")
	for _, name := range d.order {
		b.WriteString(fmt.Sprintf("  %q;\n", name))
		for _, dep := range d.nodes[name].deps {
			b.WriteString(fmt.Sprintf("  %q -> %q;\n", dep, name))
		}
	}
	b.WriteString("}\n")
	return b.String()
}
