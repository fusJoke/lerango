package main

import (
	"flag"
	"fmt"
)

var cliName   = flag.String("name", "nick", "input your name")
var cliAge    = flag.Int("age", 28, "input your age")
var cliGender = flag.String("gender", "male", "input your gender")

var cliFlag int

func Init(){
	flag.IntVar(&cliFlag, "flagName", 1234, "just for demo")
}

func main() {
	Init()
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NFlag(); i ++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("flagname=", cliFlag)
}

