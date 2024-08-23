package main

import (
	"fmt"
	"judoApi/config"
	"log"
)

func main() {
	fmt.Println("Démarrage de l'api")
	// Create logging File
	GetLogger()

	// Load the configuration file
	cfg := config.Cfg{}
	err := config.LoadConfig("config.yaml", &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Connect to Database
	config.ConnectDatabase()

	// Initialize router api
	router := InitializeRouter(&cfg)

	// Start http server
	fmt.Println("Serveur en écoute sur localhost:8080")
	log.Fatal(router.Run("localhost:8080"))
}
