package main

import (
	"context"
	"fmt"
	"time"
)

func step1(ctx context.Context) (context.Context,context.CancelFunc){
	Value := ctx.Value("1")
	fmt.Println(Value)
	select {
	case <-time.After(5*time.Second):
		fmt.Println("5 second pass")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
	default:
	}
	ctx2,cancel:= context.WithTimeout(ctx, 6*time.Second)
	go step2(ctx2)
	return ctx2, cancel
}

func step2(ctx context.Context)  {
	fmt.Println("step2 is running")
	select {
	case<-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
	default:
	}
	fmt.Println("step2 is done")
}

func main(){
	ctx := context.Background()
	ctx1, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(2*time.Second)
		cancel()
	}()

	_, cancel2 := step1(ctx1)

	go func() {
		time.Sleep(3*time.Second)
		cancel2()
	}()


}
