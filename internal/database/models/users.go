package models

import (
	"github.com/cybernuki/Service_Order_System/internal/tools"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// SchemaUser - defines a User model of the database
type SchemaUser struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index;not null;primary_key"`
	Password  string `gorm:"type:varchar(100);not null"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
}

// Create - creates a new Television regist
func (user *SchemaUser) Create() error {
	var hashedPassword string
	err := emptyFields(*user)

	if err != nil {
		return err
	}
	hashedPassword, err = HashPassword(user.Password)
	if err != nil {
		return tools.NewError("Signin not possible")
	}
	user.Password = hashedPassword
	user.ID, _ = GetUserIDByEmail(user.Email)
	isNew := Db.NewRecord(user)

	if !isNew {
		return tools.NewError("User is duplicated")
	}
	err = Db.Create(user).Error
	if err != nil {
		return tools.NewError("Signin not possible")
	}
	return nil
}

// Authenticate - search the user and compare the given password
func (user *SchemaUser) Authenticate() bool {
	var found SchemaUser
	err := Db.Find(&found, "email = ?", user.Email).Error
	if err != nil {
		return false
	}
	return CheckPasswordHash(user.Password, found.Password)
}

//GetUserIDByEmail check if a user exists in database by given email
func GetUserIDByEmail(email string) (uint, error) {
	var found SchemaUser
	err := Db.Find(&found, "email = ?", email).Error
	if err != nil {
		return 0, err
	}
	return found.ID, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func emptyFields(user SchemaUser) error {
	if user.Email == "" {
		return tools.NewError("email field cannot be empty")
	}
	if user.Password == "" {
		return tools.NewError("password field cannot be empty")
	}
	if user.FirstName == "" {
		return tools.NewError("fist_name field cannot be empty")
	}
	if user.LastName == "" {
		return tools.NewError("LastName field cannot be empty")
	}
	return nil
}
