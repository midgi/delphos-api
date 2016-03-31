package models

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// User interface
type User struct {
	name     string
	email    string
	password string
}

// NewUser Create a new User
func NewUser(name string, email string, password string) (user *User) {
	return &User{
		name:     name,
		email:    email,
		password: encrypt(password),
	}
}

// Password get user password
func (user *User) Password() string {
	return user.password
}

func encrypt(text string) string {
	key, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	hashedPassword, _ := scrypt.Key([]byte(text), []byte(key), 16384, 8, 1, 32)
	return string(hashedPassword)
}
