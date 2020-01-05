package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/satoshiyamamoto/go-gin-gorm-todos/config"
	"github.com/satoshiyamamoto/go-gin-gorm-todos/models"
	"github.com/satoshiyamamoto/go-gin-gorm-todos/routes"
)

var err error

func main() {
	// Creating a connection to the database
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	// setup routes
	r := routes.SetupRouter()

	// running
	r.Run()
}
