package gameutils

import "time"

type MandatoryGameInfo struct {
	Name       string     `json:"name"`
	Difficulty string     `json:"difficulty"`
	Board      [][]string `json:"board"`
}

type GameInfo struct {
  MandatoryGameInfo
	Id         string     `json:"uuid"`
  State      string     `json:"gameState"`
	Created    time.Time  `json:"createdAt"`
	Updated    time.Time  `json:"updatedAt"`
}
