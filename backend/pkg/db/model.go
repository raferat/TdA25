package db

import "time"

type GameBase struct {
	Name       string     `json:"name,omitempty"`
	Board      [][]string `json:"board,omitempty"`
	Difficulty string     `json:"difficulty,omitempty"`
}

type Game struct {
	GameBase

	Uuid      string    `json:"uuid,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	GameState string    `json:"gameState,omitempty"`
}
