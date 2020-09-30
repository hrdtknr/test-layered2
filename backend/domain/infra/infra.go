package infra

import (
	"fmt"
	"log"
	"github.com/test-layered2/backend/cmd/config"
	"github.com/test-layered2/backend/domain/repo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"

)

func InfraTest(){
	var t model.Todo
	t.Id = 0
	t.Name = "a"
	t.Todo = "bb"
	fmt.Println(t)
}

func GetTodo() {
	var engine *xorm.Engine

	var err error
	driverName, pass := config.Config()
	engine, err = xorm.NewEngine(driverName, pass)
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