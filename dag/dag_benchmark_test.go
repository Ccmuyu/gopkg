package dag

import (
	"context"
	"strconv"
	"testing"
)

func noopBench(_ context.Context, _ Results) (any, error) { return nil, nil }

// buildChain 构建一条 n 节点的线性依赖链
func buildChain(n int) *DAG {
	d := New()
	prev := ""
	for i := 0; i < n; i++ {
		name := strconv.Itoa(i)
		if prev == "" {
			d.Add(name, nil, noopBench)
		} else {
			d.Add(name, []string{prev}, noopBench)
		}
		prev = name
	}
	return d
}

// buildFanOut 构建一个 1 根 -> n 叶 的扇出图
func buildFanOut(n int) *DAG {
	d := New().Add("root", nil, noopBench)
	for i := 0; i < n; i++ {
		d.Add(strconv.Itoa(i), []string{"root"}, noopBench)
	}
	return d
}

func BenchmarkTopoOrder(b *testing.B) {
	d := buildChain(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = d.TopoOrder()
	}
}

func BenchmarkRunChain(b *testing.B) {
	d := buildChain(50)
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = d.Run(ctx)
	}
}

func BenchmarkRunFanOut(b *testing.B) {
	d := buildFanOut(50)
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = d.Run(ctx)
	}
}
