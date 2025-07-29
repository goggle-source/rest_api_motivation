package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model `json:"-"`
	User       string `json:"user"`
	Email      string `json:"email"`
	Category   string `json:"category"`
	Text       string `json:"text"`
}
