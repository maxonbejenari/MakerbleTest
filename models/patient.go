package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Details string `json:"details"`
}
