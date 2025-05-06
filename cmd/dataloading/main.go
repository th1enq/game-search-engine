package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/one-piece-search-engine/internal/model"
)

func main() {
	file, err := os.Open("./dataset/games_dataset.csv")
	if err != nil {
		log.Fatalf("failed to open dataset: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading CSV: %v", err)
	}

	// cfg := config.Load()
	// db := db.LoadDB(cfg)
	// es := elastic.Load(cfg)

	// gameService := service.NewGameService(db, es)
	// gameHandler := handler.NewGameHandler(gameService)

	for _, row := range records[1:] {
		release, err := strconv.Atoi(row[3])
		if err != nil {
			fmt.Printf("error formating realase: %v\n", err)
		}

		rating, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			fmt.Printf("error formating rating: %v\n", err)
		}
		game := model.Game{
			Name:     row[0],
			Genre:    row[1],
			Platform: row[2],
			Release:  release,
			Rating:   rating,
		}
		fmt.Println(game)
		// gameHandler.CreateGame
	}
}
