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

    // CORS configuration
    corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with your frontend's origin
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowCredentials: true, // If you are sending credentials like cookies
    })

    // Apply CORS middleware
    handler := corsHandler.Handler(r)

    // Define routes
    r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
    r.HandleFunc("/generate-pdf", handlers.GeneratePDFHandler).Methods("GET")


    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
