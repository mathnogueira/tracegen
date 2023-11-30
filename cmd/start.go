package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "generate a trace",
	Run: func(cmd *cobra.Command, args []string) {
		run(context.Background())
	},
}
