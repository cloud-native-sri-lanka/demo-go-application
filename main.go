package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
    // Home endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK) // 200 OK
        fmt.Fprintf(w, "Hello, welcome to our demo application v1 !")
    })

    // Crash endpoint
    http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
        fmt.Println("Exiting with status code 1...")
        os.Exit(1)
    })

    // Health check endpoint
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK) // 200 OK
        fmt.Fprintf(w, "Status: Healthy")
    })

    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}