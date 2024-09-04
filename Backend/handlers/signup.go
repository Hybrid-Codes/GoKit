package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/Hybrid-Codes/GoKit/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type SignupInput struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var input SignupInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: string(hashedPassword),
    }

    if err := models.DB.Create(&user).Error; err != nil {
        if err == gorm.ErrDuplicatedKey {
            http.Error(w, "Email already exists", http.StatusConflict)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
