package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/one-piece-search-engine/internal/db"
	"github.com/one-piece-search-engine/internal/elastic"
	"github.com/one-piece-search-engine/internal/model"
	"github.com/one-piece-search-engine/proto"
)

type GameService struct {
	db *db.DB
	es *elastic.ElasticSearch
}

func NewGameService(db *db.DB, es *elastic.ElasticSearch) *GameService {
	return &GameService{
		db: db,
		es: es,
	}
}

type ESResponse struct {
	Hits struct {
		Hits []struct {
			Source model.Game `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (s *GameService) SearchByName(ctx context.Context, name string) (*proto.Game, error) {
	query := map[string]interface{}{
		"match": map[string]interface{}{
			"name": name,
		},
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(query); err != nil {
		return nil, fmt.Errorf("failed encode query: %w", err)
	}

	res, err := s.es.Search(
		s.es.Search.WithContext(ctx),
		s.es.Search.WithIndex("games"),
		s.es.Search.WithBody(buf),
		s.es.Search.WithSize(1),
	)
	if err != nil {
		return nil, fmt.Errorf("failed search: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error elastich search: %s %s", res.Status(), string(body))
	}

	var resp ESResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("error parse json: %w", err)
	}
	if len(resp.Hits.Hits) == 0 {
		return nil, fmt.Errorf("not found")
	}
	source := resp.Hits.Hits[0].Source

	game := &proto.Game{
		Name:     source.Name,
		Genre:    source.Genre,
		Platform: source.Platform,
		Release:  int32(source.Release),
		Rating:   float32(source.Rating),
	}
	return game, nil
}

// CreateGame creates a new game in PostgreSQL only
func (s *GameService) CreateGame(ctx context.Context, protoGame *proto.Game) (*proto.Game, error) {
	game := &model.Game{
		Name:     protoGame.Name,
		Genre:    protoGame.Genre,
		Platform: protoGame.Platform,
		Release:  int(protoGame.Release),
		Rating:   float64(protoGame.Rating),
	}

	if err := s.db.Create(game).Error; err != nil {
		return nil, fmt.Errorf("failed to insert game into PostgreSQL: %w", err)
	}
	fmt.Printf("Inserted game updated_at: %v\n", game.UpdatedAt)
	return game.ProtoGame(), nil
}
