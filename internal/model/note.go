package model

import "github.com/google/uuid"

type Note struct {
	ID    uuid.UUID
	Title string
	Body  string
}

func NewNote(title, body string) *Note {
	return &Note{
		ID:    uuid.New(),
		Title: title,
		Body:  body,
	}
}
