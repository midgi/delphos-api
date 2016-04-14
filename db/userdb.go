package db

import (
   . "github.com/migdi/delphos-api/models"
	"errors"
)

type Userdb interface {
	Add(x User) User
	Contains(x User) bool
	Size() int
	
	RetrieveByEmail(x string) (User, error)
	
	
}

type userdb struct {
	users []User
}


func NewUserdb() *userdb {
	return &userdb{
	}
}


//FUNCTIONS
var id int

func (s *userdb) Add(x User) User  {
	id++
	x.SetId(id)
	
	s.users = append(s.users, x)
	
	return x
}

func (s *userdb) Size() int {
	return len(s.users)
}



func (s *userdb) Contains(x User) bool {
	for _, user := range s.users {
		if user == x {
			return true
		}
	}
	return false
}



func (s *userdb) RetrieveByEmail(x string) (User, error) {
	var nilUser User
	for _, user := range s.users {
		if user.Email() == x {
			return user, nil
		}
	}
	return nilUser,  errors.New("User not found")
}
