package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func Config() (driverName string, dataSourceName string){
	err:= godotenv.Load()
	if err != nil {
		fmt.Println("err", err)
	}
	driverName = os.Getenv("DRIVER_NAME")
	user := os.Getenv("USER_NAME")
	pass := os.Getenv("PASS")
	tcp := os.Getenv("TCP")
	dbName := os.Getenv("DB_NAME")
	dataSourceName = user+":"+pass+"@"+tcp+"/"+dbName
	return driverName, dataSourceName
}