package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context){
	select {
	case <- time.After(5*time.Second):
		fmt.Println("5 seconds pass")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	ctx.Value('1')
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	doSomething(ctx)
}