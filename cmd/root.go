package cmd

import "github.com/spf13/cobra"

var (
	collectorEndpoint string
	services          int
	minSpans          int
	traceState        []string
	insecure          bool
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
	RootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "", false, "define if http or https will be used")
	RootCmd.PersistentFlags().IntVarP(&services, "services", "s", 1, "number of services generating traces")
	RootCmd.PersistentFlags().IntVarP(&minSpans, "spans", "n", 10, "minimum number of spans")
	RootCmd.PersistentFlags().StringSliceVar(&traceState, "tracestate", []string{}, "list of tracestates to be included. First item is the name, second is the value")

	RootCmd.AddCommand(StartCmd)
	RootCmd.AddCommand(ServeCmd)
}
