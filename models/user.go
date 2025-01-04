package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null;" json:"password"`
}

func (self *User) Create() (*User, error) {
	result := db.Create(self)
	err := result.Error
	if err != nil {
		return &User{}, err
	}
	return self, nil
}

func (self *User) BeforeCreate(_ *gorm.DB) error {
	// Turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(self.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	self.Password = string(hashedPassword)

	// Remove spaces in username
	self.Username = html.EscapeString(strings.TrimSpace(self.Username))

	return nil

}
