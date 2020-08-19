package models

import (
	"math/rand"

	"github.com/cybernuki/Service_Order_System/internal/tools"
	"github.com/jinzhu/gorm"
)

// SchemaOrder - Defines an Order model of the database
type SchemaOrder struct {
	gorm.Model
	Description  string `gorm:"type:text;default:''"`
	Status       string `gorm:"type:enum('Created','Pending','Done');not null"`
	Calification int    `gorm:"type:tinyint;default:1"`
	Feedback     string `gorm:"type:tinytext;default:''"`
	User         SchemaUser
	Technician   SchemaTechnician
	Tv           SchemaTelevision
}

// Create - creates a new Television regist
func (order *SchemaOrder) Create(user SchemaUser, tv SchemaTelevision) error {
	if order.Description == "" {
		return tools.NewError("Description field cannot be empty")
	}
	order.Status = "Created"
	order.Calification = 1
	order.Feedback = ""
	order.User = user
	order.Tv = tv
	technician, err := getTechnicianRandom()

	if err != nil {
		return err
	}
	order.Technician = technician

	err = Db.Model(&SchemaOrder{}).Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func getTechnicianRandom() (SchemaTechnician, error) {
	technician := SchemaTechnician{}
	allTechnicians, err := technician.All()
	if err != nil {
		return SchemaTechnician{}, tools.NewError("Couldn't get technicians")
	}
	if len(allTechnicians) == 0 {
		return SchemaTechnician{}, tools.NewError("Not enought technicians")
	}
	technician = *allTechnicians[rand.Intn(len(allTechnicians))]
	return technician, nil
}
