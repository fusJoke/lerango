package main

import "fmt"

func main() {
	h := hello()

	if h == nil {
		fmt.Println("h is nil")
	} else {
		fmt.Println("h not nil")
	}

	h2 := hello
	if h2 == nil {
		fmt.Println("h2 is nil")
	} else {
		fmt.Println("h2 not nil")
	}

	i := GetValue()
	//只有接口类型才可以使用类型选择
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

func hello() []string {
	return nil
}

func GetValue() int {
	return 1
}
