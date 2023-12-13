package models

import (
	"gorm.io/datatypes"
)

type Tournament struct {
	ID              int             `json:"id" gorm:"primary_key"`
	Nom             string          `json:"nom"`
	Date_Debut      datatypes.Date	`json:"date_debut"`
	Date_Fin        datatypes.Date	`json:"date_fin"`
}