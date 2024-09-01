package main

import (
    "log"
    "net/http"
    "github.com/Hybrid-Codes/GoKit/models"
    "github.com/gorilla/mux"
)

func main() {
    models.ConnectDatabase()

    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
