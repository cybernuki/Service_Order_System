package models

import "github.com/jinzhu/gorm"

// SchemaTechnician - defines a Technician model of the database
type SchemaTechnician struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
}
