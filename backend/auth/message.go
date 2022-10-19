package auth

import "time"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}
type LastLoginRequest struct {
	Username  string    `json:"username" binding:"required"`
	LastLogin time.Time `json:"lastlogin" binding:"required"`
}
