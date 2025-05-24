package main

     import (
         "fmt"
         "log"
         "net/http"

         "github.com/prometheus/client_golang/prometheus"
         "github.com/prometheus/client_golang/prometheus/promhttp"
     )

     const version = "1.1.0" // Добавляем версию для отладки

     var (
         requestCounter = prometheus.NewCounter(
             prometheus.CounterOpts{
                 Name: "http_requests_total",
                 Help: "Total number of HTTP requests",
             },
         )
     )

     func init() {
         prometheus.MustRegister(requestCounter)
     }

     func handler(w http.ResponseWriter, r *http.Request) {
         requestCounter.Inc()
         log.Printf("Received request from %s", r.RemoteAddr)
         fmt.Fprintf(w, "Hello, World! Version: %s", version)
     }

     func metricsHandler() http.Handler {
         return promhttp.Handler()
     }

     func main() {
         http.HandleFunc("/metrics", metricsHandler)
         http.HandleFunc("/", handler)
         log.Printf("Server starting on port 8080, version %s...", version)
         http.ListenAndServe(":8080", nil)
     }