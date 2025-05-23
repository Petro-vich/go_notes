package noteiface

import (
	"go_notes/internal/model"
)

type NoteRepository interface {
	Create(notes *model.Note) error
	GetAll() ([]model.Note, error)
	GetById(id string) (*model.Note, error)
	Delete(id string) error
}
