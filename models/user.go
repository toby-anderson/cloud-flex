package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/toby-anderson/cloud-flex/utils/token"
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

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string,error) {
	var err error

	user := User{}

	err = db.Model(User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
