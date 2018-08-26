package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Colstuwjx/stock-exporter/src/exporter/prom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/html"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

	// prom metrics
	demoQueriesCounter = prom.New().WithCounter("demo_queries", "demo http query counter", []string{
		"name",
		"uri",
		"service_addr",
	})
)

func main() {
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/visit", func(w http.ResponseWriter, r *http.Request) {
		demoQueriesCounter.Incr("demo_queries", r.URL.Path, "127.0.0.1:8080")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
