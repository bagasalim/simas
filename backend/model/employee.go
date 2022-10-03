package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Username string `gorm:"type:varchar(191);unique" json:"username,omitempty"`
	Password string `gorm:"size:255" json:"password,omitempty"`
	Role     int    `json:"role,omitempty"`
}
