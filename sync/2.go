package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "golang.org/x/sync/errgroup"


func main() {
	g, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	serverOut := make(chan struct{})

	mux.HandleFunc("shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut<- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		case <-serverOut:
			log.Println("server will out")
		}
		timeOutCtx, cancel := context.WithTimeout(context.Background(),3*time.Second)
		defer cancel()
		log.Println("shutting down server")
		return server.Shutdown(timeOutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <- ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgourp exiting: %v\n", g.Wait())
}