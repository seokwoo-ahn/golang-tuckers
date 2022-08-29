package libs

import (
	"context"
	"fmt"
	"sync"
)

func Context() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "value", 8)
	ctx = context.WithValue(ctx, "index", 2)
	go contextTest(ctx, &wg)
	cancel()
	wg.Wait()
}

func contextTest(ctx context.Context, wg *sync.WaitGroup) {
	if v := ctx.Value("value"); v != nil {
		value := v.(int)
		fmt.Println("value:", value)
	}
	if i := ctx.Value("index"); i != nil {
		index := i.(int)
		fmt.Println("index:", index)
	}
	wg.Done()
}
