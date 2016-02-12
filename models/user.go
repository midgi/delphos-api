package models

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// User interface
type User struct {
	Name     string
	Email    string
	Password string
}

// NewUser Create a new User
func NewUser(name string, email string, password string) (user *User) {
	return &User{
		Name:     name,
		Email:    email,
		Password: encrypt(password),
	}
}

// GetPassword get user password
func (user *User) GetPassword() string {
	return user.Password
}

func encrypt(text string) string {
	key, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	hashedPassword, _ := scrypt.Key([]byte(text), []byte(key), 16384, 8, 1, 32)
	return string(hashedPassword)
}
