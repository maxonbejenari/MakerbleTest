package models

import "gorm.io/gorm"

type UserRole string

const (
	Receptionist UserRole = "receptionist"
	Doctor       UserRole = "doctor"
)

type User struct {
	gorm.Model
	Username string   `gorm:"uniqueIndex" json:"username"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}
