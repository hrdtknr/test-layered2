package main

import (
	"fmt"
	"log"
	"github.com/test-layered2/backend/cmd/config"
	"github.com/test-layered2/backend/domain/infra"
	"github.com/test-layered2/backend/domain/service"
	"github.com/test-layered2/backend/app/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main () {
	fmt.Println("main")
	driverName, dataSourceName := config.Config()
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer engine.Close()

	engine.ShowSQL(true)
	engine.SetMapper(core.GonicMapper{})

	e:= echo.New()


	todoRepo := infra.NewTodo(engine)
	todoService := service.NewTodo(todoRepo)
	log.Println(todoService)
	//h := handler.NewHandler(todoService)

}