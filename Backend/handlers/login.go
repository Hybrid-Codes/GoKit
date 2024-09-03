package handlers

import (
    "encoding/json"
    "net/http"
	"github.com/Hybrid-Codes/GoKit/models"
	"github.com/Hybrid-Codes/GoKit/utils"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var input LoginInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user models.User
    if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
    http.Error(w, "Error generating token", http.StatusInternalServerError)
    return
	}

	response := map[string]interface{}{
		"user":  user,
		"token": token,
	}

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode(response)
}
