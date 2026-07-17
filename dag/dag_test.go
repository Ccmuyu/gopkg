package dag

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	gtest "github.com/Ccmuyu/gopkg/test"
)

func noop(_ context.Context, _ Results) (any, error) { return nil, nil }

func TestRunOrder(t *testing.T) {
	var mu sync.Mutex
	var seq []string
	record := func(name string) Task {
		return func(_ context.Context, _ Results) (any, error) {
			mu.Lock()
			seq = append(seq, name)
			mu.Unlock()
			return name, nil
		}
	}

	d := New().
		Add("a", nil, record("a")).
		Add("b", []string{"a"}, record("b")).
		Add("c", []string{"a"}, record("c")).
		Add("d", []string{"b", "c"}, record("d"))

	res, err := d.Run(context.Background())
	gtest.AssertNoError(t, err)
	gtest.AssertEqual(t, res["d"], "d")

	// a 必须最先，d 必须最后
	gtest.AssertEqual(t, seq[0], "a")
	gtest.AssertEqual(t, seq[len(seq)-1], "d")
	gtest.AssertEqual(t, len(seq), 4)
}

func TestDataFlow(t *testing.T) {
	d := New().
		Add("x", nil, func(_ context.Context, _ Results) (any, error) {
			return 10, nil
		}).
		Add("y", nil, func(_ context.Context, _ Results) (any, error) {
			return 20, nil
		}).
		Add("sum", []string{"x", "y"}, func(_ context.Context, deps Results) (any, error) {
			return deps["x"].(int) + deps["y"].(int), nil
		})

	res, err := d.Run(context.Background())
	gtest.AssertNoError(t, err)
	gtest.AssertEqual(t, res["sum"], 30)
}

func TestCycleDetection(t *testing.T) {
	d := New().
		Add("a", []string{"c"}, noop).
		Add("b", []string{"a"}, noop).
		Add("c", []string{"b"}, noop)

	_, err := d.TopoOrder()
	gtest.AssertError(t, err)
	gtest.AssertError(t, d.Validate())

	_, runErr := d.Run(context.Background())
	gtest.AssertError(t, runErr)
}

func TestSelfCycle(t *testing.T) {
	d := New().Add("a", []string{"a"}, noop)
	gtest.AssertError(t, d.Validate())
}

func TestMissingDependency(t *testing.T) {
	d := New().Add("a", []string{"ghost"}, noop)
	err := d.Validate()
	gtest.AssertError(t, err)
	gtest.AssertTrue(t, gtest.ContainsStr(err.Error(), "ghost"))
}

func TestDuplicateTask(t *testing.T) {
	d := New().
		Add("a", nil, noop).
		Add("a", nil, noop)
	gtest.AssertError(t, d.Validate())
}

func TestNilFunc(t *testing.T) {
	d := New().Add("a", nil, nil)
	gtest.AssertError(t, d.Validate())
}

func TestEmptyName(t *testing.T) {
	d := New().Add("", nil, noop)
	gtest.AssertError(t, d.Validate())
}

func TestEmptyDAG(t *testing.T) {
	res, err := New().Run(context.Background())
	gtest.AssertNoError(t, err)
	gtest.AssertEqual(t, len(res), 0)
}

func TestFailFast(t *testing.T) {
	sentinel := errors.New("boom")
	var downstreamRan atomic.Bool

	d := New().
		Add("a", nil, func(_ context.Context, _ Results) (any, error) {
			return nil, sentinel
		}).
		Add("b", []string{"a"}, func(_ context.Context, _ Results) (any, error) {
			downstreamRan.Store(true)
			return nil, nil
		})

	_, err := d.Run(context.Background())
	gtest.AssertError(t, err)
	gtest.AssertTrue(t, errors.Is(err, sentinel))
	gtest.AssertTrue(t, !downstreamRan.Load())
}

func TestContinueOnError(t *testing.T) {
	sentinel := errors.New("boom")
	var cRan atomic.Bool

	d := New().
		Add("a", nil, func(_ context.Context, _ Results) (any, error) {
			return nil, sentinel
		}).
		Add("b", []string{"a"}, func(_ context.Context, _ Results) (any, error) {
			return "should-skip", nil
		}).
		Add("c", nil, func(_ context.Context, _ Results) (any, error) {
			cRan.Store(true)
			return "ok", nil
		})

	res, err := d.Run(context.Background(), WithContinueOnError())
	gtest.AssertError(t, err)
	gtest.AssertTrue(t, errors.Is(err, sentinel))
	// c 与 a 无依赖关系，应继续执行
	gtest.AssertTrue(t, cRan.Load())
	gtest.AssertEqual(t, res["c"], "ok")
	// b 依赖失败的 a，应被跳过（无输出）
	_, ok := res["b"]
	gtest.AssertTrue(t, !ok)
}

func TestMaxConcurrency(t *testing.T) {
	var running, maxSeen int64
	makeTask := func() Task {
		return func(ctx context.Context, _ Results) (any, error) {
			cur := atomic.AddInt64(&running, 1)
			for {
				old := atomic.LoadInt64(&maxSeen)
				if cur <= old || atomic.CompareAndSwapInt64(&maxSeen, old, cur) {
					break
				}
			}
			time.Sleep(20 * time.Millisecond)
			atomic.AddInt64(&running, -1)
			return nil, nil
		}
	}

	d := New()
	for _, name := range []string{"a", "b", "c", "d", "e", "f"} {
		d.Add(name, nil, makeTask())
	}

	_, err := d.Run(context.Background(), WithMaxConcurrency(2))
	gtest.AssertNoError(t, err)
	gtest.AssertTrue(t, atomic.LoadInt64(&maxSeen) <= 2)
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	d := New().
		Add("a", nil, func(ctx context.Context, _ Results) (any, error) {
			cancel()
			<-ctx.Done()
			return nil, ctx.Err()
		}).
		Add("b", []string{"a"}, func(_ context.Context, _ Results) (any, error) {
			return "ran", nil
		})

	_, err := d.Run(ctx)
	gtest.AssertError(t, err)
}

func TestPreCanceledContextReturnsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	d := New().Add("a", nil, noop)
	results, err := d.Run(ctx)
	gtest.AssertTrue(t, errors.Is(err, context.Canceled))
	gtest.AssertEqual(t, len(results), 0)
}

func TestExternalCancellationIsReturnedWhenTaskSucceeds(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	d := New().Add("a", nil, func(context.Context, Results) (any, error) {
		cancel()
		return "done", nil
	})

	results, err := d.Run(ctx)
	gtest.AssertTrue(t, errors.Is(err, context.Canceled))
	gtest.AssertEqual(t, results["a"], "done")
}

func TestTopoOrder(t *testing.T) {
	d := New().
		Add("a", nil, noop).
		Add("b", []string{"a"}, noop).
		Add("c", []string{"b"}, noop)

	order, err := d.TopoOrder()
	gtest.AssertNoError(t, err)
	gtest.AssertSliceEqual(t, order, []string{"a", "b", "c"})
}

func TestDOT(t *testing.T) {
	d := New().
		Add("a", nil, noop).
		Add("b", []string{"a"}, noop)
	dot := d.DOT()
	gtest.AssertTrue(t, gtest.ContainsStr(dot, "digraph dag"))
	gtest.AssertTrue(t, gtest.ContainsStr(dot, `"a" -> "b"`))
}

func TestMetadata(t *testing.T) {
	d := New().
		Add("a", nil, noop).
		Add("b", []string{"a"}, noop)
	gtest.AssertEqual(t, d.Len(), 2)
	gtest.AssertSliceEqual(t, d.TaskNames(), []string{"a", "b"})
}
