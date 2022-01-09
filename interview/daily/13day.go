package main

import "fmt"

func main() {
	i := 5
	defer hello(i)
	j := i + 10
	fmt.Printf("j = %d\n",j)
}

func hello(i int)  {
	fmt.Printf("i = %d\n",i)
}

//hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5