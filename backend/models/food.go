package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name       string `json:"name"`
	Quantity   string `json:"quantity"`
	Restaurant string `json:"restaurant"`
	Status     string `json:"status" gorm:"default:Available"`
}
