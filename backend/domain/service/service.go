package service

import (
	"log"
	"github.com/test-layered2/backend/domain/repo"
	"github.com/test-layered2/backend/domain/repo/model"
)

type Todo interface{
	Todos() ([]model.Todo, error)
}

type todo struct {
	repo repo.Todo
}

func NewTodo(r repo.Todo) Todo{
	log.Println("service NewTodo")
	log.Println("service r repo.Todo\n",r)
	return &todo{repo: r}
}

func (t *todo) Todos() ([]model.Todo, error){
	return t.repo.Todos()
}