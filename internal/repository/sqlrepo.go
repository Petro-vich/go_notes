package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go_notes/internal/model"
	"go_notes/pkg/noteiface"
	"log"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(pathDb string) (noteiface.NoteRepository, error) {
	dbSql, err := sql.Open("sqlite3", pathDb)
	if err != nil {
		return nil, fmt.Errorf("Не удалость открыть базу данны: %w", err)
	}

	expression := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL
)`
	_, err = dbSql.Exec(expression)
	if err != nil {
		return nil, fmt.Errorf("error creating the table: %w", err)
	}

	NewRep := SQLiteRepo{
		db: dbSql,
	}

	return &NewRep, nil
}

func (r *SQLiteRepo) Create(notes *model.Note) error {
	expression := `INSERT INTO notes (title, body, user_id) VALUES (?, ?, ?)`
	_, err := r.db.Exec(expression, notes.Title, notes.Body, notes.ID)
	if err != nil {
		return fmt.Errorf("error creating the note: %w", err)
	}

	return nil
}
func (r *SQLiteRepo) GetAll() ([]*model.Note, error) {
	query := `SELECT user_id, title, body FROM notes`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting the notes: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}()

	var listNotes []*model.Note

	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Body); err != nil {
			return nil, fmt.Errorf("error getting the notes: %w", err)
		}
		listNotes = append(listNotes, &note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting the notes: %w", err)
	}

	return listNotes, nil
}
func (r *SQLiteRepo) GetByID(id string) (*model.Note, error) {
	query := `SELECT title, body, user_id FROM notes WHERE id = ?`
	row := r.db.QueryRow(query, id)

	result := &model.Note{}
	if err := row.Scan(&result.Title, &result.Body, &result.ID); err != nil {
	}

	return result, nil
}
func (r *SQLiteRepo) Delete(id string) error {
	expression := `DELETE FROM notes WHERE id = ?`
	_, err := r.db.Exec(expression, id)
	if err != nil {
		return fmt.Errorf("error deleting the note: %w", err)
	}

	return nil
}
