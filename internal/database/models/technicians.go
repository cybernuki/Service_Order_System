package models

import (
	"github.com/cybernuki/Service_Order_System/internal/tools"
	"github.com/jinzhu/gorm"
)

// SchemaTechnician - defines a Technician model of the database
type SchemaTechnician struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index;not null;primary_key"`
	Password  string `gorm:"type:varchar(100);unique_index"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
}

// Create - creates a new Television regist
func (technician *SchemaTechnician) Create() error {
	var hashedPassword string
	err := emptyFieldsTechnician(*technician)

	if err != nil {
		return err
	}
	hashedPassword, err = HashPassword(technician.Password)
	if err != nil {
		return tools.NewError("Signin not possible")
	}
	technician.Password = hashedPassword
	technician.ID, _ = GetTechnicianIDByEmail(technician.Email)
	isNew := Db.NewRecord(technician)

	if !isNew {
		return tools.NewError("Technician is duplicated")
	}
	err = Db.Create(technician).Error
	if err != nil {
		return tools.NewError("Signin not possible")
	}
	return nil
}

// Authenticate - search the user and compare the given password
func (technician *SchemaTechnician) Authenticate() bool {
	var found SchemaTechnician
	err := Db.Find(&found, "email = ?", technician.Email).Error
	if err != nil {
		return false
	}
	return CheckPasswordHash(technician.Password, found.Password)
}

//GetTechnicianIDByEmail check if a user exists in database by given email
func GetTechnicianIDByEmail(email string) (uint, error) {
	var found SchemaTechnician
	err := Db.Find(&found, "email = ?", email).Error
	if err != nil {
		return 0, err
	}
	return found.ID, nil
}

func emptyFieldsTechnician(technician SchemaTechnician) error {
	if technician.Email == "" {
		return tools.NewError("email field cannot be empty")
	}
	if technician.Password == "" {
		return tools.NewError("password field cannot be empty")
	}
	if technician.FirstName == "" {
		return tools.NewError("fist_name field cannot be empty")
	}
	if technician.LastName == "" {
		return tools.NewError("LastName field cannot be empty")
	}
	return nil
}
