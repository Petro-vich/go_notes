/*
	di - Dependency Injection(Внедрение зависимостей)

*/

package di

import (
	"fmt"
	"go_notes/internal/repository"
	"go_notes/pkg/noteiface"
)

const (
	RepoTypeJSON = "JSON"
)

/*
Инициализация репозитория. Возвращаемся интерфейс NoteRepository. Почему интерфейс, а не конкретную структуру?
Баз данных можем быть несколько и что бы не писать под каждую базу данных свою функцию с нужным типом возвращаемых
данных, мы возвращаем интерфейс, который в сам определяет возвращаемую структуру.
Например: у SQL и у JSON один интерфейс NoteRepository который исполняет одни и теже методы
*/

func InitRepository(typeRep string, pathRep string) (noteiface.NoteRepository, error) {
	switch typeRep {
	case RepoTypeJSON:
		return repository.NewJSONRepo(pathRep), nil
	default:
		return nil, fmt.Errorf("di: invalid repository type %q", typeRep)
	}
}
