package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SignInRequest struct {
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SignInRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, SignInResponse{Error: fmt.Sprintf("Failed to read request: %v", err)})
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &req); err != nil {
		writeJSON(w, SignInResponse{Error: fmt.Sprintf("Invalid JSON format: %v", err)})
		return
	}

	expectedPassword := os.Getenv("TODO_PASSWORD")
	if expectedPassword == "" || req.Password != expectedPassword {
		writeJSON(w, SignInResponse{Error: "Invalid password"})
		return
	}

	token, err := generateToken(req.Password)
	if err != nil {
		writeJSON(w, SignInResponse{Error: fmt.Sprintf("Failed to generate token: %v", err)})
		return
	}

	writeJSON(w, SignInResponse{Token: token})
}
