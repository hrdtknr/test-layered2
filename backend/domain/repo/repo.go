package repo

import (
	"github.com/test-layered2/backend/domain/repo/model"
)

type Todo interface {
	Todos() ([]model.Todo, error)
}