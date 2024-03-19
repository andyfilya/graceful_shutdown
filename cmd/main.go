package main

import (
	"context"
	"github.com/andyfilya/graceful_shutdown/pkg/server"
	errgroup "golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	server := server.InitServer(ctx, "8081")
	// создаём канал для прослушки системных сигналов в горутине
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

		<-ch // слушаем данный канал до тех пор, пока не получим какой-то сигнал
		cancel()
	}()

	gr, grCtx := errgroup.WithContext(ctx)
	gr.Go(func() error {
		return server.ListenAndServe()
	})

	gr.Go(func() error {
		<-grCtx.Done()
		return server.Shutdown(context.Background())
	})

	if err := gr.Wait(); err != nil {
		log.Fatal(err)
	}
}
