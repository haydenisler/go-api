package cmd

import (
	"context"
	"fmt"

	"github.com/haydenisler/api/internal/api"
	"github.com/spf13/cobra"
)

func APICmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Args:  cobra.ExactArgs(0),
		Short: "Runs the REST API",
		RunE: func(cmd *cobra.Command, args []string) error {
			port := 8000

			api := api.NewAPI(ctx)
			srv := api.Server(port)

			go func() { _ = srv.ListenAndServe() }()

			fmt.Printf("started api on port %d\n", port)

			// Blocks until a value is passed on the done ch
			<-ctx.Done()

			_ = srv.Shutdown(ctx)

			return nil
		},
	}

	return cmd
}
