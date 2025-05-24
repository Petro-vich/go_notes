/*
	di - Dependency Injection(Внедрение зависимостей)

*/

package di

import (
	"fmt"
	"go_notes/internal/repository"
	"go_notes/pkg/noteiface"
	"os"
)

const (
	DbTypeJson   = "JSON"
	DbTypeSQLite = "SQLITE"
)

/*
Инициализация репозитория. Возвращаемся интерфейс NoteRepository. Почему интерфейс, а не конкретную структуру?
Баз данных можем быть несколько и что бы не писать под каждую базу данных свою функцию с нужным типом возвращаемых
данных, мы возвращаем интерфейс, который в сам определяет возвращаемую структуру.
Например: у SQL и у JSON один интерфейс NoteRepository который исполняет одни и теже методы
*/

func InitRepository() (noteiface.NoteRepository, error) {
	repoType := os.Getenv("REPO_TYPE")
	repoPath := os.Getenv("REPO_PATH")

	switch repoType {
	case DbTypeJson:
		return repository.NewJSONRepo(repoPath), nil
	case DbTypeSQLite:
		rep, err := repository.NewSQLiteRepo(repoPath)
		if err != nil {
			return nil, err
		}
		return rep, nil
	default:
		return nil, fmt.Errorf("di: invalid repository type %q", repoType)
	}
}
