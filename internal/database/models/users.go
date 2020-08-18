package models

import (
	"github.com/jinzhu/gorm"
)

// SchemaUser - defines a User model of the database
type SchemaUser struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
}
