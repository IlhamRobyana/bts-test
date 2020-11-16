package entity

import (
	"time"
)

// RegisterRequest is used as the request body on register
type RegisterRequest struct {
	User User
}

// LoginRequest is used as the request body on login
type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// LoginResponse is used as the response on login
type LoginResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"Username"`
}

// ValidateTokenResponse is used as the response on ValidateToken
type ValidateTokenResponse struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Timestamp time.Time `json:"timestamp"`
}
