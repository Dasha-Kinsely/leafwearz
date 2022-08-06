package entities

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Username string `gorm:"column:username;unique"`
	Email string `gorm:"column:email;unique"`
	Password string `gorm:"column:password;size:100;notNull"`
	IsAdmin bool `gorm:"column:admin;notNull"`
}