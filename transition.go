package main

import (
	_ "github.com/lib/pq"
)

//SongSuggestion object for returning songlist
type Transition struct {
	FromSong  SongDetailed `json:"FromSong,omitempty"`
	ToSong    SongDetailed `json:"ToSong,omitempty"`
	Occasions int64        `json:"Occasions,omitempty"`
	Score     float64      `json:"Score,omitempty"`
}
