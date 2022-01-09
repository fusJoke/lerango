package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	later := time.Now().Add(time.Minute)
	fmt.Println(later)
	fmt.Println(later.Sub(now))
	fmt.Println(later.Equal(now))
	fmt.Println(later.Before(now))
	fmt.Println(later.After(now))
}