package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	_, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", respBytes)
}
