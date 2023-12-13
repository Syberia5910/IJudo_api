package main

import (
	"judoApi/config"
	"log"
	"fmt"
  )
  
  func main() {
	fmt.Println("Démarrage de l'api")
	// Create logging File
	GetLogger()

	// Connect to Database
	config.ConnectDatabase()

	// Initialize router api
	router := InitializeRouter()

	// Start http server
	fmt.Println("Serveur en écoute sur localhost:8080")
	log.Fatal(router.Run("localhost:8080"))
  }