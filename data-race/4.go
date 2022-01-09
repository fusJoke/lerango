package main

import (
	"fmt"
	"os"
	"time"
)

type Watchdog struct{ last int64}

func(w *Watchdog) KeepAlive(){
	w.last = time.Now().UnixNano()
}

func (w *Watchdog) Start(){
	go func() {
		for {
			time.Sleep(time.Second)
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. DYing")
				os.Exit(1)
			}
		}
	}()
}
