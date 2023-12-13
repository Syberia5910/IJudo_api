package controllers

import (
  "log"
  "net/http"
  "judoApi/models"
  "judoApi/config"
  "github.com/gin-gonic/gin"
)

// GET /clubs
// Get all clubs
func GetClubs(c *gin.Context) {
  var clubs []models.Club
  config.DB.Find(&clubs)

  c.JSON(http.StatusOK, clubs)
}

// GET /club/:id
// Find a club
func GetClub(c *gin.Context) {
	var club models.Club
	if err := config.DB.Where("ID = ?", c.Param("id")).First(&club).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, club)
}

// POST /club
// Create new club
func CreateClub(c *gin.Context) {
  log.Println("Début création club --")
  // Validate input
  var input models.Club
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Problème de structure des éléments envoyés (" + err.Error() + ")"})
    return
	}

  // Create Club
  club := models.Club{Nom: input.Nom}
  config.DB.Create(&club)
  log.Println("Fin création club -- OK 200")
  c.JSON(http.StatusOK, club)
}