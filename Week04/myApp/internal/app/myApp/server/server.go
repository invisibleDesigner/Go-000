package server

import (
	"context"
	"errors"
	"fmt"
	"go_playground/Go-000/Week04/myApp/api/rest"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func NewServer()  {
	engine := rest.SetUp()
	srv := http.Server{
		Addr:              ":8081",
		Handler:           engine,
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		go func() {
			err := srv.ListenAndServe()
			if err != nil {
				fmt.Println(err)
				cancel()
			}
		}()
		select {
		case <- ctx.Done():
			srv.Shutdown(ctx)
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
