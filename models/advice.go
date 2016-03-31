package models

import "time"

//Advice Interface
type Advice struct {
	id              string
	content         string
	adviceRequester User
	createdAt       time.Time
}

// NewAdvice new advice
func NewAdvice(id string, content string, adviceRequester User) (advice *Advice) {
	return &Advice{
		id:              id,
		content:         content,
		adviceRequester: adviceRequester,
		createdAt:       time.Now(),
	}
}

// ID returns advice id
func (advice *Advice) ID() string {
	return advice.id
}

// Content returns content
func (advice *Advice) Content() string {
	return advice.content
}
