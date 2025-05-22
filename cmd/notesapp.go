package main

import (
	"fmt"
	"github.com/lpernett/godotenv"
	"go_notes/internal/di"
	"go_notes/internal/model"
	"log"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("config/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	typeRepo := os.Getenv("REPO_TYPE")
	pathRepo := os.Getenv("REPO_PATH")
	rep, err := di.InitRepository(typeRepo, pathRepo)
	if err != nil {
		log.Fatal(err)
	}

	nt := model.NewNote("Тест", "Удалить эту заметку")
	if err := rep.Create(nt); err != nil {
		log.Fatal(err)
	}

	ntAll, err := rep.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, nt := range ntAll {
		fmt.Println(nt.Title)
		fmt.Println(nt.Body)
		fmt.Printf("\n")
	}

}
