package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/propagation"
)

var serverPort int

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a HTTP server for generating a trace",
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
			ctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			fmt.Println("Start tracing generator")
			go run(ctx)
			fmt.Println("Trace generating finished")
		})

		fmt.Println("Server started")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
	},
}

func init() {
	ServeCmd.Flags().IntVarP(&serverPort, "port", "p", 8088, "port to run the http server")
}
