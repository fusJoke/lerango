package main

import (
	"fmt"
	"time"
)

type chanTest struct {
	done  chan struct{}
}

func (c *chanTest) Done() <-chan struct{} {
	c.done = make(chan struct{})
	d := c.done
	return d
}

func handler(ct *chanTest){

	done := ct.Done()

	select {
	case <-done:
		fmt.Println("1111")
	}
}
func main() {
	ct := chanTest{}
	go handler(&ct)

	time.Sleep(2*time.Second)
}