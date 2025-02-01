package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/haydenisler/api/internal/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	result := cmd.Execute(ctx)

	os.Exit(result)
}
