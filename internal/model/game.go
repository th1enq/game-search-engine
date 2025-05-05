package model

type Game struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	Genre    string  `json:"genre"`
	Platform string  `json:"platform"`
	Release  int     `json:"realease"`
	Rating   float64 `json:"rating"`
}
