package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func initAuth() {
	password := os.Getenv("TODO_PASSWORD")
	if password != "" {
		jwtSecret = []byte(password)
	}
}

func generateToken(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	hashStr := hex.EncodeToString(hash[:])

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"password_hash": hashStr,
		"exp":           time.Now().Add(8 * time.Hour).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func validateToken(tokenString string) bool {
	if len(jwtSecret) == 0 {
		return true
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	password := os.Getenv("TODO_PASSWORD")
	hash := sha256.Sum256([]byte(password))
	currentHash := hex.EncodeToString(hash[:])

	storedHash, ok := claims["password_hash"].(string)
	if !ok || storedHash != currentHash {
		return false
	}

	return true
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := os.Getenv("TODO_PASSWORD")
		if password != "" {
			var tokenStr string

			cookie, err := r.Cookie("token")
			if err == nil {
				tokenStr = cookie.Value
			}

			if !validateToken(tokenStr) {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}
		}

		next(w, r)
	}
}
