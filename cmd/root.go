package cmd

import "github.com/spf13/cobra"

var (
	collectorEndpoint string
	services          int
	minSpans          int
)

var RootCmd = &cobra.Command{
	Use:   "tracegen",
	Short: "generate a complex trace",
}

func Execute() {
	RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&collectorEndpoint, "collector", "c", "localhost:4317", "grpc endpoint for your collector")
	RootCmd.PersistentFlags().IntVarP(&services, "services", "s", 1, "number of services generating traces")
	RootCmd.PersistentFlags().IntVarP(&minSpans, "spans", "n", 10, "minimum number of spans")

	RootCmd.AddCommand(StartCmd)
	RootCmd.AddCommand(ServeCmd)
}
