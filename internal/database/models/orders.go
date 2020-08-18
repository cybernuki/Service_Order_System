package models

import (
	"github.com/jinzhu/gorm"
)

// Order - Defines an Order model of the database
type Order struct {
	gorm.Model
	Status       string     `gorm:"type:enum('Created','Pending','Done');not null"`
	Calification int        `gorm:"type:tinyint;default:1"`
	Feedback     string     `gorm:"type:tinytext;default:''"`
	URL          string     `gorm:"type:tinytext;default:''"`
	User         User       `gorm:"foreignkey:ID"`
	Technician   Technician `gorm:"foreignkey:ID"`
	Tv           Television `gorm:"foreignkey:ID"`
}
