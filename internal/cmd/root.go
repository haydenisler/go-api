package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {
	rootCmd := &cobra.Command{
		Use:   "~ [command]",
		Short: "ToT",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(APICmd(ctx))

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
