package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Television - defines a Television model of the database
type Television struct {
	gorm.Model
	ModelTV string `gorm:"type:varchar(50);not null"`
	Brand   string `gorm:"type:varchar(50);not null"`
}

// Create - creates a new Television regist
func (tv *Television) Create() {
	newTv := tv
	if Db != nil && Db.DB().Ping() != nil {
		log.Panicln("Database session is down")
	}

	Db.Find(&newTv, "model_tv = ? AND brand = ?", tv.ModelTV, tv.Brand)
	if newTv.ID != tv.ID {
		Db.Create(tv)
	}
}
