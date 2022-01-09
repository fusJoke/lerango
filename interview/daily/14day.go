package main

func add(args ...int) int {

	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}


