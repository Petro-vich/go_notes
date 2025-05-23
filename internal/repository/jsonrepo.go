package repository

import (
	"encoding/json"
	"fmt"
	"go_notes/dto"
	"go_notes/internal/model"
	"go_notes/pkg/noteiface"
	"os"
	"sync"
)

type JSONRepo struct {
	filepath string
	mu       sync.Mutex
}

func NewJSONRepo(filepath string) noteiface.NoteRepository {
	return &JSONRepo{filepath: filepath}
}

func (r *JSONRepo) Create(note *model.Note) error {
	//Получили все структуры типа model.Note из JSON file
	existing, err := r.GetAll()
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	//Проверка, есть ли Note с тем же ID
	for _, nt := range existing {
		if nt.ID == note.ID {
			return fmt.Errorf("note already exists")
		}
	}

	//Добавляем новый NOTE
	existing = append(existing, model.Note{
		ID:    note.ID,
		Title: note.Title,
		Body:  note.Body,
	})

	//Переводим структуру Note в структуру NoteDTO
	var dtos []dto.NoteDTO
	for _, note := range existing {
		dtos = append(dtos, dto.FromEntity(note))
	}

	//Кодируем перед записью
	data, err := json.MarshalIndent(dtos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filepath, data, 0644)
}

func (r *JSONRepo) GetAll() ([]model.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	//Считываем JSON в data
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Note{}, nil
		}
		return nil, err
	}

	//Декодируем JSON из data и заполняем dtos структуру
	var dtos []dto.NoteDTO
	if err := json.Unmarshal(data, &dtos); err != nil {
		return nil, err
	}

	//Приводим dtos структуру к структуре Note
	var result []model.Note
	for _, d := range dtos {
		result = append(result, d.ToModel())
	}
	return result, err
}

func (r *JSONRepo) Delete(id string) error {

	data, err := r.GetAll()
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	var newData []model.Note
	for _, nt := range data {
		if nt.ID == id {
			continue
		}
		newData = append(newData, nt)
	}

	var dtos []dto.NoteDTO
	for _, d := range newData {
		dtos = append(dtos, dto.FromEntity(d))
	}

	result, err := json.MarshalIndent(newData, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filepath, result, 0644)
}

func (r *JSONRepo) GetById(id string) (*model.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

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
