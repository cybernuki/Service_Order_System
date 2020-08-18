package models

import (
	"log"

	"github.com/cybernuki/Service_Order_System/internal/tools"

	"github.com/jinzhu/gorm"
)

// SchemaTelevision - defines a Television model of the database
type SchemaTelevision struct {
	gorm.Model
	ModelTV string `gorm:"type:varchar(50);not null"`
	Brand   string `gorm:"type:varchar(50);not null"`
}

// Create - creates a new Television regist
func (tv *SchemaTelevision) Create() error {
	newTv := tv
	if Db != nil && Db.DB().Ping() != nil {
		log.Panicln("Database session is down")
	}
	if tv.ModelTV == "" {
		return tools.NewError("Model field cannot be empty")
	}
	if tv.Brand == "" {
		return tools.NewError("Brand field cannot be empty")
	}
	Db.Find(&newTv, "model_tv = ? AND brand = ?", tv.ModelTV, tv.Brand)
	if newTv.ID == 0 || newTv.ID != tv.ID {
		Db.Create(tv)
	} else {
		return tools.NewError("Television is duplicated")
	}
	return nil
}

// All - return all the televisions in the database
func (tv *SchemaTelevision) All() ([]*SchemaTelevision, error) {
	var televisions []*SchemaTelevision
	err := Db.Find(&televisions).Error
	if err != nil {
		log.Println(err)
	}
	return televisions, err
}
