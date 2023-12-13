package models

import (
  "gorm.io/datatypes"
)

type Category struct {
  ID                  int                 `json:"id" gorm:"primary_key"`
  Nom                 string              `json:"nom"`
  Sexe                string              `json:"sexe"`
  Date_Naissance_min  datatypes.Date      `json:"date_naissance_min"`
  Date_Naissance_max  datatypes.Date      `json:"date_naissance_max"`
  Poid_Min            int                 `json:"poid_min`
  Poid_Max            int                 `json:"poid_max`
  Date_Ouverture      datatypes.Date      `json:"date_ouverture"`
  Date_Cloture        datatypes.Date	    `json:"date_cloture"`
  TournamentID        int                 `json:"tournament_ID"`
  Tournament          Tournament          `gorm:"foreignKey:TournamentID" gorm:"references:ID"`
}