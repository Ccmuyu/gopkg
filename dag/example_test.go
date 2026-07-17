package dag_test

import (
	"context"
	"fmt"

	"github.com/Ccmuyu/gopkg/dag"
)

func ExampleDAG_Run() {
	d := dag.New().
		Add("fetch", nil, func(_ context.Context, _ dag.Results) (any, error) {
			return 7, nil
		}).
		Add("double", []string{"fetch"}, func(_ context.Context, deps dag.Results) (any, error) {
			return deps["fetch"].(int) * 2, nil
		}).
		Add("report", []string{"double"}, func(_ context.Context, deps dag.Results) (any, error) {
			return fmt.Sprintf("result=%d", deps["double"].(int)), nil
		})

	res, err := d.Run(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(res["report"])
	// Output: result=14
}

func ExampleDAG_TopoOrder() {
	noop := func(_ context.Context, _ dag.Results) (any, error) { return nil, nil }
	d := dag.New().
		Add("a", nil, noop).
		Add("b", []string{"a"}, noop).
		Add("c", []string{"a"}, noop).
		Add("d", []string{"b", "c"}, noop)

	order, _ := d.TopoOrder()
	fmt.Println(order)
	// Output: [a b c d]
}

func ExampleDAG_Validate() {
	d := dag.New().
		Add("a", []string{"b"}, func(_ context.Context, _ dag.Results) (any, error) { return nil, nil }).
		Add("b", []string{"a"}, func(_ context.Context, _ dag.Results) (any, error) { return nil, nil })

	if err := d.Validate(); err != nil {
		fmt.Println("invalid:", err != nil)
	}
	// Output: invalid: true
}
