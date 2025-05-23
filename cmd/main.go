package main

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"go_notes/internal/app"
)

func init() {
	err := godotenv.Load("/home/gaylesga/Desktop/s21.project/go_notes/config/.env")
	if err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if err := app.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
