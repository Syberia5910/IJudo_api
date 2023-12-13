package models

import (
  "gorm.io/datatypes"
)

type Judoka struct {
  ID              int             `json:"id" gorm:"primary_key"`
  Nom             string          `json:"nom"`
  Prenom          string          `json:"prenom"`
  Sexe            string          `json:"sexe"`
  Date_Naissance  datatypes.Date	`json:"date_naissance"`
  ClubID          int             `json:"club_ID" `
  Club            Club            `gorm:"foreignKey:ClubID" gorm:"references:ID"`
}