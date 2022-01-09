package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 9 ,20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s))

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Cap, cap(s))
}
