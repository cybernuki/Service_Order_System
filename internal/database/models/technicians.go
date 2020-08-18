package models

import "github.com/jinzhu/gorm"

// SchemaTechnician - defines a Technician model of the database
type SchemaTechnician struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index;not null;primary_key"`
	Password  string `gorm:"type:varchar(100);unique_index"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
}
