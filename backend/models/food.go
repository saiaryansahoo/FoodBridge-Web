package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name       string `json:"name"`
	Quantity   int `json:"quantity"`
	Restaurant string `json:"restaurant"`
	Status     string `json:"status" gorm:"default:Available"`
}
