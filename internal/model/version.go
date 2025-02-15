package model

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}
