package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var pingCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "Ping_request_count",
	Help: "Number of pings made to the endpoint",
})

func HiEndpoint(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "Hi")

}

func main() {
	prometheus.MustRegister(pingCounter)
	http.HandleFunc("/hello", HiEndpoint)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)

}
