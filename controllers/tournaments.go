package controllers

import (
  "log"
	"net/http"
	"judoApi/models"
	"judoApi/config"
	"github.com/gin-gonic/gin"
)

// GET /tournaments
// Get all tournaments
func GetTournaments(c *gin.Context) {
  var tournaments []models.Tournament
  	config.DB.Find(&tournaments)

	c.JSON(http.StatusOK, tournaments)
}

// GET /tournament/:id
// Find a tournament
func GetTournament(c *gin.Context) {
	var tournament models.Tournament
	if err := config.DB.Where("ID = ?", c.Param("id")).First(&tournament).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, tournament)
}

// POST /tournament
// Create new tournament
func CreateTournament(c *gin.Context) {
  log.Println("Début création tournoi --")
  // Validate input
  var input models.Tournament
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Problème de structure des éléments envoyés (" + err.Error() + ")"})
    return
  }

	// Create Tournament
	tournament := models.Tournament{Nom: input.Nom, Date_Debut: input.Date_Debut, Date_Fin: input.Date_Fin}
	config.DB.Create(&tournament)
	c.JSON(http.StatusOK, tournament)
	log.Println("Fin création tournoi -- OK", http.StatusOK)
}