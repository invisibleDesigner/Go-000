package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	bizMux := http.NewServeMux()
	bizMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
	})
	bizServer := http.Server{
		Addr:              ":8081",
		Handler:           bizMux,
	}
	eg.Go(func() error {
		go func() {
			err := bizServer.ListenAndServe()
			if err != nil {
				fmt.Println(err)
				cancel()
			}
		}()
		select {
		case <- ctx.Done():
			bizServer.Shutdown(ctx)
			return ctx.Err()
		}
	})
	debugMux := http.NewServeMux()
	debugMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "debug")
	})
	debugServer := http.Server{
		Addr:              ":8082",
		Handler:           debugMux,
	}
	eg.Go(func() error {
		go func() {
			err := debugServer.ListenAndServe()
			if err != nil {
				fmt.Println(err)
				cancel()
			}
		}()
		select {
		case <- ctx.Done():
			debugServer.Shutdown(ctx)
			return ctx.Err()
		}
	})
	eg.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			return errors.New(sig.String())
		case <- ctx.Done():
			return ctx.Err()
		}
	})
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

// https://github.com/go-kratos/kratos/blob/v2/app.go