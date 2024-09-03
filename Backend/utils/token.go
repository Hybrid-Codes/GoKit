package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var secretKey = []byte("your-secret-key")

func GenerateToken(email string) (string, error) {
    claims := jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}
