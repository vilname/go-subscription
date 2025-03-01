package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"subscription-back/config"
	"subscription-back/config/storage"
)

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8082
func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env: ", err.Error())
	}

	storage.InitDB()
	router := config.InitRoute()

	defer func(ctx context.Context) {
		db := storage.GetDB()
		db.Close()
	}(context.Background())

	fmt.Println("init")

	err = router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
