package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/propagation"
)

var serverPort int
var delay string

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a HTTP server for generating a trace",
	Run: func(cmd *cobra.Command, args []string) {
		delayDuration, err := time.ParseDuration(delay)
		if err != nil {
			fmt.Printf("failed to parse delay duration: %s\n", err.Error())
			os.Exit(1)
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
			ctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			fmt.Println("Start tracing generator")
			go run(ctx, WithDelay(delayDuration))
			fmt.Println("Trace generating finished")
		})

		fmt.Printf("Server started on port %d\n", serverPort)
		if delayDuration > 0 {
			fmt.Printf("All traces will be generated after %s\n", delayDuration.String())
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
	},
}

func init() {
	ServeCmd.Flags().IntVarP(&serverPort, "port", "p", 8088, "port to run the http server")
	ServeCmd.Flags().StringVarP(&delay, "delay", "d", "0ms", "time before start sending spans (100ms, 2s, 3m, etc)")
}
