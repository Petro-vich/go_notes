package repository

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go_notes/dto"
	"go_notes/internal/model"
	"os"
	"sync"
)

type JSONRepo struct {
	filepath string
	mu       sync.Mutex
}

func NewJSONRepo(filepath string) NoteRepository {
	return &JSONRepo{filepath: filepath}
}

func (r *JSONRepo) Create(note *model.Note) error {
	existing, err := r.GetAll()
	if err != nil {
		return err
	}

	for _, nt := range existing {
		if nt.ID == note.ID {
			return fmt.Errorf("note already exists")
		}
	}

	existing = append(existing, model.Note{
		ID:    note.ID,
		Title: note.Title,
		Body:  note.Body,
	})

	data, err := json.MarshalIndent(existing, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filepath, data, 0644)
}

func (r *JSONRepo) GetAll() ([]model.Note, error) {
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Note{}, nil
		}
	}
	var dtos []dto.NoteDto
	var result []model.Note
	if err := json.Unmarshal(data, &dtos); err != nil {
		return nil, err
	}

	for _, d := range dtos {
		parsedID, err := uuid.Parse(d.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, model.Note{
			ID:    parsedID,
			Title: d.Title,
			Body:  d.Body,
		})
	}
	return result, err
}

func (r *JSONRepo) Delete(id uuid.UUID) error {
	data, err := r.GetAll()
	if err != nil {
		return err
	}

	result := make([]model.Note, len(data))
	for _, nt := range data {
		if nt.ID == id {
			continue
		}
		result = append(result, nt)
	}

	return nil
}

func (r *JSONRepo) GetById(id uuid.UUID) (*model.Note, error) {
	data, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, nt := range data {
		if nt.ID == id {
			return &nt, nil
		}
	}

	return nil, fmt.Errorf("note not found")
}
