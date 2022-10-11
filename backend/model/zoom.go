package model

import "gorm.io/gorm"

type Zoom struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(30); not null;unique"`
	Email    string `json:"email" gorm:"type:varchar(100); not null"`
	Kategori string `json:"kategori" gorm:"type:varchar(50); not null"`
	Keluhan  string `json:"keluhan" gorm:"type:varchar(200); not null"`
}
