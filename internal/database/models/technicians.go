package models

import "github.com/jinzhu/gorm"

// Technician - defines a Technician model of the database
type Technician struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
}
