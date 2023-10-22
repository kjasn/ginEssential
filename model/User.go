package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(20);not null"`
	Telephone string `json:"telephone" gorm:"type:varchar(11);not null;unique"` // 手机号唯一
	Pwd       string `json:"pwd" gorm:"type:varchar(255);not null"`
}
