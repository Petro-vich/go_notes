package dto

import "go_notes/internal/model"

type NoteDTO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (d NoteDTO) ToModel() model.Note {
	return model.Note{
		ID:    d.ID,
		Title: d.Title,
		Body:  d.Body,
	}
}

func FromEntity(n model.Note) NoteDTO {
	return NoteDTO{
		ID:    n.ID,
		Title: n.Title,
		Body:  n.Body,
	}
}
