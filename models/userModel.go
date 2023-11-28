package models

import "test/utils"

type User struct {
	utils.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
