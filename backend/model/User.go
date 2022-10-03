package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" `
	Task     string `json:"task"`
	Done     bool   `json:"done"`
}
