package models

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// User interface
type User struct {
	id 		 int
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


// get the user Id
func (user *User) SetId(id int)  {
	user.id = id
}

// get the user Id
func (user *User) Id() int {
	return user.id
}

// Password get user password
func (user *User) Password() string {
	return user.password
}

// Password get user password
func (user *User) Email() string {
	return user.email
}

// get user name
func (user *User) Name() string {
	return user.name
}


func encrypt(text string) string {
	key, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	hashedPassword, _ := scrypt.Key([]byte(text), []byte(key), 16384, 8, 1, 32)
	return string(hashedPassword)
}
