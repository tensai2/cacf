package models

import "gorm.io/gorm"

type Facebook struct {
	gorm.Model
	Token string
	Name  string
	Email string
}
