package main

import (
	"net/http"
    "math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_simple_gauge",
		Help: "This is a simple gauge metric example.",
	})
	
	prometheus.MustRegister(gauge)

    http.Handle("/metrics", prometheusMiddleware(gauge, promhttp.Handler()))

	http.ListenAndServe(":8080", nil)
}

func prometheusMiddleware(gauge prometheus.Gauge, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gauge.Set(float64(rand.Intn(100))) 
		next.ServeHTTP(w, r)
	})
}
