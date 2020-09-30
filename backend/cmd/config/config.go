package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func Config() (d1 string, d2 string){
	//mainと同じディレクトリに.envを置く必要がある
	err:= godotenv.Load()
	if err != nil {
		fmt.Println("err", err)
	}
	driverName := os.Getenv("DRIVER_NAME")
	user := os.Getenv("USER_NAME")
	pass := os.Getenv("PASS")
	tcp := os.Getenv("TCP")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(driverName)
	fmt.Println(user+":"+pass+"@"+tcp+"/"+dbName)
	return driverName, user+":"+pass+"@"+tcp+"/"+dbName
}