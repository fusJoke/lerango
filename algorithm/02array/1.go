package main

import "sync"

type Array struct {
	array []int
	cap  int
	len int
	lock sync.Mutex
}

func makeArray(len, cap int)  *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}

	array := make([]int, cap, cap)
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

func append(a *Array, element int)  {
	a.lock.Lock()
	defer a.lock.Lock()

	if a.cap == a.len {
		newCap := 2*a.len

		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)
		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = element
	a.len = a.len +1
}

func main() {
	
}
