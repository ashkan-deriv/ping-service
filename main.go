package main

import (
        "log"
        "net/http"
)

func main() {
        http.HandleFunc("/ping", handlePing)
        log.Printf("Server starting on :8080")
        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatal(err)
        }
}

func handlePing(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message":"pong"}`))
}