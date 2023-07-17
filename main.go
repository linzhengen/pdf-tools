package main

import (
	"context"
	"github.com/linzhengen/pdf-tools/cmd"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if err := cmd.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
