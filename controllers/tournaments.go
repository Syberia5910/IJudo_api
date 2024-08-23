package controllers

import (
  "log"
	"net/http"
	"judoApi/models"
	"judoApi/config"
  	"gorm.io/datatypes"
	"github.com/gin-gonic/gin"
    "time"
	"fmt"
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
func CreateTournament(c *gin.Context, cfg *config.Cfg) {
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
	// Create categories
	for _, cat := range cfg.BaseCategory {
		// Convert Date_Debut to time.Time
		dateDebut := time.Time(tournament.Date_Debut)
		dateMin := time.Date(dateDebut.AddDate(- cat.Variance,0,0).Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		dateMax := time.Date(dateDebut.AddDate(- cat.Variance,0,0).AddDate(- cat.Gap,0,0).AddDate(1,0,0).Year(), time.December, 31, 23, 59, 59, 0, time.UTC)
		fmt.Println(dateMin, dateMax)
		category := models.Category{
		 	Nom: cat.Name,
			Date_Naissance_max: datatypes.Date(dateMin),
			Date_Naissance_min: datatypes.Date(dateMax),
			TournamentID: tournament.ID,
			Sexe: cat.Genre,
		}
		config.DB.Preload("Tournament").Create(&category)
	}

	c.JSON(http.StatusOK, tournament)
	log.Println("Fin création tournoi -- OK", http.StatusOK)
}