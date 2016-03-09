package models

import (
 "time"
)

//Advice Interface
type Advice struct {
	Id 		string
	Content	string
	AdviceRequester User
	CreatedAt	time.Time
}

//New Advice
func NewAdvice(id string, content string, adviceRequester User) (advice *Advice) {
	return &Advice{
		Id:					id,
		Content:				content,
		AdviceRequester: 	adviceRequester,
		CreatedAt: 			time.Now(),
	}
	
}

func (advice *Advice) GetId() string {
	return advice.Id
}

func (advice *Advice) GetContent() string {
	return advice.Content
}
