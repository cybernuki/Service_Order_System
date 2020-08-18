package models

import (
	"github.com/jinzhu/gorm"
)

// SchemaOrder - Defines an Order model of the database
type SchemaOrder struct {
	gorm.Model
	Description  string           `gorm:"type:text;default:''"`
	Status       string           `gorm:"type:enum('Created','Pending','Done');not null"`
	Calification int              `gorm:"type:tinyint;default:1"`
	Feedback     string           `gorm:"type:tinytext;default:''"`
	URL          string           `gorm:"type:tinytext;default:''"`
	User         SchemaUser       `gorm:"foreignkey:ID"`
	Technician   SchemaTechnician `gorm:"foreignkey:ID"`
	Tv           SchemaTelevision `gorm:"foreignkey:ID"`
}
