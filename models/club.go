package models

type Club struct {
	ID   		int    `json:"id" gorm:"primary_key"`
	Nom  		string `json:"nom"`
	NomCourt	string `json:"nom_court`
}