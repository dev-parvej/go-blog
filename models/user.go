package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Posts    []Post
}
