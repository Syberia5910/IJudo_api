package controllers

import (
	"log"
	"net/http"
	"judoApi/models"
	"judoApi/config"
	"github.com/gin-gonic/gin"
)


// GET /categories/:tournamentID
// Get all category
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Preload("Tournament").Where("tournament_id = ?", c.Param("tournamentID")).Find(&categories).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, categories)
}


// GET /categories/:tournamentID/open
// Get all category
func GetOpenCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Preload("Tournament").Where("tournament_id = ? AND date_ouverture != ? AND date_cloture == ?", c.Param("tournamentID"), "0001-01-01 00:00:00+00:00", "0001-01-01 00:00:00+00:00").Find(&categories).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, categories)
}


// GET /category/:id
// Find a category
func GetCategory(c *gin.Context) {
  var category models.Category
  if err := config.DB.Preload("Tournament").Where("ID = ?", c.Param("id")).First(&category).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, category)
}


//POST /categorie
// Create new categorie
func CreateCategory(c *gin.Context) {
	log.Println("Début création categorie --")
	// Validate input
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problème de structure des éléments envoyés (" + err.Error() + ")"})
		return
	}

	// Create Category
	category := models.Category{ Nom: input.Nom, Date_Ouverture: input.Date_Ouverture, 
    Date_Cloture: input.Date_Cloture, Date_Naissance_max: input.Date_Naissance_max,
    Date_Naissance_min: input.Date_Naissance_min, TournamentID: input.TournamentID,
    Sexe: input.Sexe }
  config.DB.Preload("Tournament").Create(&category)
  c.JSON(http.StatusOK, category)
  log.Println("Fin création category -- OK", http.StatusOK)
}