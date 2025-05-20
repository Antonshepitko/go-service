package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received request from %s", r.RemoteAddr)
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}