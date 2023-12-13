package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"judoApi/config"
	"judoApi/models"
)


// GET /judokas
// Get all judokas
func GetJudokas(c *gin.Context) {
	log.Println("Début GetAllJudokas")
	var judokas []models.Judoka
	config.DB.Preload("Club").Find(&judokas)

	c.JSON(http.StatusOK, judokas)
	log.Println("Fin GetAllJudokas --", http.StatusOK)
}


// GET /judoka/:id
// Find a judoka
func GetJudoka(c *gin.Context) {
	var judoka models.Judoka
	if err := config.DB.Preload("Club").Where("ID = ?", c.Param("id")).First(&judoka).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, judoka)
}


// POST /judoka
// Create new judoka
func CreateJudoka(c *gin.Context) {
	log.Println("Début création judoka --")
	
	// Validate input
	var input models.Judoka

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "CreateJudoka - Problème de structure des éléments envoyés (" + err.Error() + ")"})
		return 
	}
	
	// Create Judoka
	judoka := models.Judoka{Nom: input.Nom, Prenom: input.Prenom, 
		Date_Naissance: input.Date_Naissance, ClubID: input.ClubID,
		Sexe: input.Sexe}
	config.DB.Preload("Club").Create(&judoka)

	log.Println("Fin création judoka --", http.StatusOK)
	c.JSON(http.StatusOK, judoka)
}
