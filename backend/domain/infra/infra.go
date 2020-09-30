package infra

import (
	"log"
	//"github.com/test-layered2/backend/cmd/config"
	"github.com/test-layered2/backend/domain/repo"
	"github.com/test-layered2/backend/domain/repo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"xorm.io/core"

)

type todo struct {
	engine *xorm.Engine
}

func NewTodo(engine *xorm.Engine) repo.Todo {
	log.Println("infra NewTodo")
	log.Println("engine\n", engine)
	return &todo{engine: engine}
}

func (t *todo) Todos()([]model.Todo, error){
	todos := []model.Todo{}
	if err := t.engine.Find(&todos); err != nil {
		log.Println("error of GetTodoList\n", err)
		return nil, err
	}
	log.Println("infra todos", todos)
	return todos, nil
}







/*
// ここ以下は問題なく動作
func GetTodo() {
	var engine *xorm.Engine

	var err error
	driverName, dataSourceName := config.Config()
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer engine.Close()

	engine.ShowSQL(true)
	engine.SetMapper(core.GonicMapper{})


	var todoList []model.Todo
	err = engine.Find(&todoList)

	if err != nil {
		log.Println(err)
	}
	log.Println(todoList)

}
*/