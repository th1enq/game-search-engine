package handler

import (
	"context"
	"fmt"

	"github.com/one-piece-search-engine/internal/service"
	"github.com/one-piece-search-engine/proto"
)

type GameHandler struct {
	gameService *service.GameService
}

func NewGameHandler(gameService *service.GameService) *GameHandler {
	return &GameHandler{
		gameService: gameService,
	}
}

func (g *GameHandler) CreateGame(ctx context.Context, req *proto.CreateGameRequest) (*proto.GamesResponse, error) {
	if req.Game == nil {
		return nil, fmt.Errorf("game information is required")
	}

	// Call the service to create the game
	createdGame, err := g.gameService.CreateGame(ctx, req.Game)
	if err != nil {
		return nil, fmt.Errorf("failed to create game: %w", err)
	}

	// Return the created game in the response
	return &proto.GamesResponse{
		Game: []*proto.Game{createdGame},
	}, nil
}

func (g *GameHandler) UpdateGame(ctx context.Context, req *proto.UpdateGameRequest) (*proto.GamesResponse, error) {

	return nil, nil
}

func (g *GameHandler) DeleteGame(ctx context.Context, req *proto.CreateGameRequest) (*proto.GamesResponse, error) {

	return nil, nil
}

func (g *GameHandler) SearchByName(ctx context.Context, req *proto.SearchByNameRequest) (*proto.GameResponse, error) {
	if req.Game == nil {
		return nil, fmt.Errorf("game information is required")
	}

	name := req.Game.Name
	if name == "" {
		return nil, fmt.Errorf("game name is required")
	}

	game, err := g.gameService.SearchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &proto.GameResponse{
		Game: game,
	}, nil
}

func (g *GameHandler) SearchByGenre(ctx context.Context, req *proto.SearchByGenreRequest) (*proto.GamesResponse, error) {

	return nil, nil
}

func (g *GameHandler) SearchByPlatform(ctx context.Context, req *proto.SearchByPlatformRequest) (*proto.GamesResponse, error) {

	return nil, nil
}

func (g *GameHandler) SearchByReleaseRange(ctx context.Context, req *proto.SearchByReleaseRangeRequest) (*proto.GamesResponse, error) {

	return nil, nil
}

func (g *GameHandler) SearchByRating(ctx context.Context, req *proto.SearchByRatingRequest) (*proto.GamesResponse, error) {

	return nil, nil
}

func (g *GameHandler) AdvancedSearch(ctx context.Context, req *proto.AdvancedSearchRequest) (*proto.GamesResponse, error) {

	return nil, nil
}
