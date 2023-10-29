package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, str string) {
	fmt.Printf("Running function foo for %s\n", str)

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Running for 500 milisecond with message %s\n", str)
	case <-ctx.Done():
		fmt.Printf("Canceling foo with message %s\n", str)
		return
	}

	fmt.Printf("Finishing foo with message %s\n", str)
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go foo(ctx, "call-1")
	go foo(ctx, "call-2")

	time.Sleep(100 * time.Millisecond)

	cancel()

	fmt.Printf("Back to the main routine!!\n")

}
