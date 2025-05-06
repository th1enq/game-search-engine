package model

import (
	"github.com/one-piece-search-engine/proto"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name     string  `json:"name" gorm:"not null"`
	Genre    string  `json:"genre"`
	Platform string  `json:"platform"`
	Release  int     `json:"realease"`
	Rating   float64 `json:"rating"`
}

func (g *Game) ProtoGame() *proto.Game {
	pg := proto.Game{
		Name:     g.Name,
		Genre:    g.Genre,
		Platform: g.Platform,
		Release:  int32(g.Release),
		Rating:   float32(g.Rating),
	}
	return &pg
}
