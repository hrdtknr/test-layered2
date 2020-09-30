package repo

import (
	"test-layered/backend/domain/repo/model"
)

type IUserRepository interface {
	GetTodoList() ([]model.Todo, error)
}