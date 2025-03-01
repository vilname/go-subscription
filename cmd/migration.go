package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"subscription-back/config"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env: ", err.Error())
	}

	config.InitMigrationDB()

	fmt.Println("миграции отработали")
}
