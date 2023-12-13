package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"judoApi/models"
	"judoApi/config"
)

// POST /inscription
// Register a Judoka to a tournament (categorie)
func CreateRegistration(c *gin.Context) {
	// Validate input
	var inputRegistration models.Registration
	if err := c.ShouldBindJSON(&inputRegistration); err != nil {
	 	log.Println(err)
	 	c.JSON(http.StatusBadRequest, gin.H{"error": "CreateRegistration -- Problème de structure des éléments envoyés (" + err.Error() + ")"})
	 	return 
	}

	// Create registration
	log.Println(inputRegistration.Poid)
	registration := models.Registration{Poid: inputRegistration.Poid,
		Ceinture: inputRegistration.Ceinture, JudokaID: inputRegistration.JudokaID}
	config.DB.Preload("Judoka").Preload("Club").Create(&registration)
	
	c.JSON(http.StatusOK, registration)
	log.Println("Fin inscription judoka --", http.StatusOK)
}

// GET /registrations
// Find all registrations
func GetRegistrationsPerTournament(c *gin.Context) {
	var registrations []models.Registration
	if err := config.DB.Preload("Judoka").Find(&registrations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, registrations)
}