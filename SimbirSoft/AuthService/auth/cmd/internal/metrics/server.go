package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Listen(port string) error {
	//use separated ServeMux to prevent handling on the global Mux
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting metrics server at %s/metrics", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to serve metrics: %v", err)
		return err
	}

	return nil
}
