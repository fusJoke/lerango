package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler start")
	ctx := r.Context()
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("5 seconds pass")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
	fmt.Println("handler ends")

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
