package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main(){
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i <5; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	//main一边向i写，一边起goroutine读取i进行print。出现了数据竞争
	wg.Wait()
}
