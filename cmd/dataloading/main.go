package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/one-piece-search-engine/config"
	"github.com/one-piece-search-engine/internal/db"
	"github.com/one-piece-search-engine/internal/elastic"
	"github.com/one-piece-search-engine/internal/handler"
	"github.com/one-piece-search-engine/internal/model"
	"github.com/one-piece-search-engine/internal/service"
	"github.com/one-piece-search-engine/proto"
	"google.golang.org/grpc"
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

	cfg := config.Load()
	db := db.LoadDB(cfg)
	es := elastic.Load(cfg)

	gameService := service.NewGameService(db, es)
	gameHandler := handler.NewGameHandler(gameService)

	s := grpc.NewServer()

	proto.RegisterGamesServer(s, gameHandler)

	for i, row := range records[1:] {
		release, err := strconv.Atoi(row[3])
		if err != nil {
			fmt.Printf("error formatting release for row %d: %v\n", i+1, err)
			continue // Skip this row and continue with next
		}

		rating, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			fmt.Printf("error formatting rating for row %d: %v\n", i+1, err)
			continue // Skip this row and continue with next
		}

		game := model.Game{
			Name:     row[0],
			Genre:    row[1],
			Platform: row[2],
			Release:  release,
			Rating:   rating,
		}

		// Convert model.Game to proto.Game
		protoGame := game.ProtoGame()

		// Save game to database
		_, err = gameService.CreateGame(context.Background(), protoGame)
		if err != nil {
			fmt.Printf("error adding game '%s' to database: %v\n", game.Name, err)
			continue
		}

		fmt.Printf("Successfully added game: %s\n", game.Name)
	}

	fmt.Println("Data loading completed!")
}
