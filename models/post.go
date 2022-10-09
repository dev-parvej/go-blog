package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body" gorm:"type=text"`
	UserID int    `json:"user_id"`
	User   User
}
