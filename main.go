package main

     import (
         "fmt"
         "log"
         "net/http"

         "github.com/prometheus/client_golang/prometheus"
         "github.com/prometheus/client_golang/prometheus/promhttp"
     )

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
         fmt.Fprintf(w, "Hello, World!")
     }

     func main() {
         http.HandleFunc("/", handler)
         http.Handle("/metrics", promhttp.Handler())
         log.Println("Server starting on port 8080...")
         http.ListenAndServe(":8080", nil)
     }