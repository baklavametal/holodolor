package main

import (
	"fmt"
	"log"

	"github.com/baklavametal/holodolor/backend/pkg/api"
	"github.com/baklavametal/holodolor/backend/pkg/config"
	"github.com/baklavametal/holodolor/backend/pkg/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Read configuration
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error reading configuration: %s", err)
		return
	}

	// Initialize the database
	db, err := gorm.Open("postgres", conf.DatabaseURL)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Auto-migrate all models
	db.AutoMigrate(
		&model.User{},
		&model.Event{},
		&model.EventType{},
		&model.AdditionalProp{},
		&model.EventTypeProp{},
	)

	// Initialize Gin
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true // This is for demonstration; in production, set this to your frontend's origin
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")

	r.Use(cors.New(corsConfig))

	// Setup routes
	api.SetupRoutes(r, db)

	// Start the server
	r.Run(":8088")
}
