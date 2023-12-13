package models

type Registration struct {
	ID			int 		`json:"id" gorm:"primary_key"`
	Poid		float32		`json:"poid"`
	Ceinture	string		`json:"ceinture"`
	JudokaID	int			`json:"judoka_ID"`
	Judoka		Judoka		`gorm:"foreignKey:JudokaID" gorm:"references:ID"`
	CategoryID	int			`json:"category_ID"`
	Category	Category	`gorm:"foreignKey:CategoryID" gorm:"references:ID"`
}