package main

import (
	
	"log"
	

	"github.com/Himneesh-Kalra/makerble-coding-assessment/api"
	"github.com/Himneesh-Kalra/makerble-coding-assessment/db"
	"github.com/Himneesh-Kalra/makerble-coding-assessment/models"
	

	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system env.")
	}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	
	err = db.AutoMigrate(&models.User{}, &models.Patient{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	
	server := api.NewApiServer(db)
	log.Println("Server starting on :8080")
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}


