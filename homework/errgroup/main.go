package main

import (
	"context"
	"fmt"
	"github.com/go-errors/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	//模拟http server关闭
	serverClose := make(chan struct{})
	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		serverClose <- struct{}{}
	})

	server := http.Server{
		Addr: ":8888",
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler: mux,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
			case <-ctx.Done():
				log.Println("Errgroup exit...")
			case <-serverClose:
				log.Println("Server will close...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		defer cancel()

		log.Println("Server is closing...")
		return server.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-quit:
				return errors.Errorf("Linux signal: %v", sig)
		}
	})

	fmt.Printf("Errgroup: %+v\n", g.Wait())
}