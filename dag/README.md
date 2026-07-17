# dag

有向无环图（DAG）任务编排工具：声明任务及其依赖，执行时自动拓扑排序、环检测，并在满足依赖约束下最大化并行。支持 `context` 取消、并发上限、快速失败 / 继续执行，以及任务间数据传递。

## 安装

```go
import "github.com/Ccmuyu/gopkg/dag"
```

## 快速开始

```go
d := dag.New().
    Add("fetch", nil, func(ctx context.Context, deps dag.Results) (any, error) {
        return 7, nil
    }).
    Add("double", []string{"fetch"}, func(ctx context.Context, deps dag.Results) (any, error) {
        return deps["fetch"].(int) * 2, nil
    }).
    Add("report", []string{"double"}, func(ctx context.Context, deps dag.Results) (any, error) {
        return fmt.Sprintf("result=%d", deps["double"].(int)), nil
    })

res, err := d.Run(context.Background())
if err != nil {
    log.Fatal(err)
}
fmt.Println(res["report"]) // result=14
```

## 核心概念

- **任务（Task）**：由唯一 `name`、依赖列表 `deps` 和执行函数 `fn` 组成。
- **数据传递**：`fn` 的 `deps dag.Results` 参数包含该任务【直接依赖】的输出；`fn` 的返回值成为自己的输出，供下游读取。
- **并行**：所有依赖已满足的任务会并发执行，无需手动管理 goroutine。

```go
type Task func(ctx context.Context, deps Results) (any, error)
type Results map[string]any
```

## API

| 方法 | 说明 |
|---|---|
| `New()` | 创建空图 |
| `Add(name, deps, fn) *DAG` | 添加任务，支持链式调用 |
| `Validate() error` | 校验依赖缺失、重名、空名、nil 函数与环 |
| `TopoOrder() ([]string, error)` | 返回稳定的拓扑排序序列 |
| `Run(ctx, opts...) (Results, error)` | 并发执行，返回结果与聚合错误 |
| `DOT() string` | 导出 Graphviz DOT，便于可视化 |
| `Len()` / `TaskNames()` | 任务数量 / 任务名列表 |

## 执行选项

```go
// 限制同时运行的任务数（<=0 不限制）
d.Run(ctx, dag.WithMaxConcurrency(4))

// 出错继续：仅跳过失败任务的下游，其余互不依赖的任务照常执行
d.Run(ctx, dag.WithContinueOnError())
```

## 错误处理

- **默认（快速失败）**：任一任务返回错误即取消 `ctx`，尚未开始的任务被跳过；`Run` 返回该错误。
- **`WithContinueOnError`**：尽力执行，跳过失败任务的下游；所有失败被聚合返回。
- 返回的错误基于 `errors.MultiError`，实现了 `Unwrap() []error`，可用 `errors.Is` / `errors.As` 定位具体任务错误。
- 无论成功与否，`Run` 都会返回已完成任务的部分结果。

```go
res, err := d.Run(ctx)
if errors.Is(err, ErrUpstream) {
    // 命中特定上游错误
}
```

## 环 / 非法依赖检测

`Validate`、`TopoOrder`、`Run` 都会检测：

- 依赖了不存在的任务
- 重名 / 空名 / nil 函数
- 存在环（包括自依赖），报错会列出涉及的节点

## 注意事项

- `Run` 采用协作式取消：快速失败时正在运行的任务不会被强杀，请在 `fn` 中尊重 `ctx` 以尽早退出。
- 传给 `fn` 的 `deps` 是当次执行的只读快照副本，可安全并发读取。
