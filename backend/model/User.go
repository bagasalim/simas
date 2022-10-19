package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username" gorm:"type:varchar(30); not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(100); not null"`
	Name      string    `json:"name" gorm:"type:varchar(50); not null"`
	Role      int8      `json:"role"  `
	LastLogin time.Time `json:"lastlogin"`
}
