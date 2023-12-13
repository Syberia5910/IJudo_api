package models

type Tatami struct {
	ID					int					`json:"id" gorm:"primary_key"`
	Nom					string				`json:"tatami"`
	TournamentID        int                 `json:"tournament_ID"`
	Tournament          Tournament          `gorm:"foreignKey:TournamentID" gorm:"references:ID"`
}