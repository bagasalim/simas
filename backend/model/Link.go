package model

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Link_Value string `json:"linkvalue" gorm:"type:varchar(256);"`
	Link_Type  string `json:"linktype" gorm:"type:varchar(100);"`
	Updated_By string `json:"updatedby" gorm:"type:varchar(30);"`
}
