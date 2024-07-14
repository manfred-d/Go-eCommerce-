package main

import (
	"backend/go_backend/config"
	"backend/go_backend/database"
	"backend/go_backend/models"
	"backend/go_backend/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := database.ConnectDB(&config)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// database migration
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate to database", err)
	}

	fmt.Println("Successful connected to database")

	router := gin.Default()

	routes.SetupRoutes(router, db)

	log.Printf("Server starting ... %s", config.ServerAddr)

	if err := router.Run(config.ServerAddr); err != nil {
		log.Fatal("Failed to connect to server", err)
	}

}
