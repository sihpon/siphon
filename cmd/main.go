package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/siphon/siphon/cmd/config"
	"github.com/siphon/siphon/cmd/server"
	siphonoperator "github.com/siphon/siphon/cmd/siphon-operator"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	_, err := config.NewConfigFromFlags()
	if err != nil {
		panic(err)
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return siphonoperator.Start()
	})
	eg.Go(func() error {
		return server.Start()
	})

	quit := make(chan os.Signal, 1)
	defer signal.Stop(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		fmt.Println("ctx.Done() received")
		break
	case sig := <-quit:
		fmt.Printf("Signal received: %v\n", sig)
		break
	}

	//
	// server.Close()
	// if err = eg.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 	panic(err)
	// }
}
