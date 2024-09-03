package main

import (
    "log"
    "net/http"
    "github.com/Hybrid-Codes/GoKit/models"
    "github.com/Hybrid-Codes/GoKit/handlers"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    models.ConnectDatabase()

    r := mux.NewRouter()

    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"}, // Replace with your frontend's origin
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Authorization", "Content-Type"},
    })

    r.Use(corsHandler.Handler)

    r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
