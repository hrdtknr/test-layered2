package main

import (
	"fmt"
	"github.com/test-layered2/backend/cmd/config"
	"github.com/test-layered2/backend/domain/infra"
)

func main () {
	fmt.Println("main")
	config.Config()
	infra.InfraTest()
	infra.GetTodo()
}