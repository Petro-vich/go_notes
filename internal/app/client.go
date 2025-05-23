package app

import (
	"errors"
	"fmt"
	"go_notes/internal/di"
	"go_notes/internal/model"
	"os"
	"strings"
)

func Run(args []string) error {
	if len(args) == 0 {
		return errors.New("no command provided (use: create, list, delete)")
	}

	cmd := strings.ToLower(args[0])
	repoType := getenv("REPO_TYPE", "JSON")
	repoPath := getenv("REPO_PATH", "tasks.json")

	repo, err := di.InitRepository(repoType, repoPath)
	if err != nil {
		return err
	}

	switch cmd {
	case "create":
		if len(args) < 3 {
			return errors.New("usage: create <title> <body>")
		}
		n := model.NewNote(args[1], args[2])
		if err := repo.Create(n); err != nil {
			return err
		}
		fmt.Println("Note created with ID:", n.ID)

	case "list":
		notes, err := repo.GetAll()
		if err != nil {
			return err
		}
		for _, note := range notes {
			fmt.Printf("ID: %s\nTitle: %s\nBody: %s\n\n", note.ID, note.Title, note.Body)
		}

	case "delete":
		if len(args) < 2 {
			return errors.New("usage: delete <note_id>")
		}
		if err := repo.Delete(args[1]); err != nil {
			return err
		}
		fmt.Println("Note deleted.")

	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}

	return nil
}

func getenv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
