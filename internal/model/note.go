package model

import (
	"github.com/google/uuid"
)

type Note struct {
	ID    string
	Title string
	Body  string
}

func NewNote(title, body string) *Note {
	generateId := uuid.NewString()
	return &Note{
		ID:    generateId,
		Title: title,
		Body:  body,
	}
}
