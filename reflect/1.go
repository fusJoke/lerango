package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct {

}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	var animal Animal
	var cat Cat
	//interface变量可以存储任意实现了该接口类型的变量
	//interface类型的变量在存储某个变量时同时保存变量类型和变量值
	animal = cat
	fmt.Println(animal.Speak())
}