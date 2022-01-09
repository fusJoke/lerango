package main

import "fmt"

func main(){
	i := [5]int{1,2,3,4,5}
	j := i[3:4:4]
	fmt.Println(j[0])

	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {	//compilation error
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
