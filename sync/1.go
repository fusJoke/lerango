package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func hello(){
	defer wg.Done()
	fmt.Println("112323")
}

func main(){
	wg.Add(1)
	go hello()
	fmt.Println("main goroutine is running")
	wg.Wait()
}


