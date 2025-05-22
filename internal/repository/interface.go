package repository

import (
	"github.com/google/uuid"
	"go_notes/internal/model"
)

type NoteRepository interface {
	Create(notes *model.Note) error
	GetAll() ([]model.Note, error)
	GetById(id uuid.UUID) (*model.Note, error)
	Delete(id uuid.UUID) error
}
