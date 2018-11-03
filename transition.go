package main

import (
	_ "github.com/lib/pq"
)

//TransitionDetailed ...
type TransitionDetailed struct {
	FromSong  SongDetailed `json:"FromSong,omitempty"`
	ToSong    SongDetailed `json:"ToSong,omitempty"`
	Occasions int64        `json:"Occasions,omitempty"`
	Score     float64      `json:"Score,omitempty"`
}

//Transition ...
type Transition struct {
	ToSong Song    `json:"ToSong,omitempty"`
	Score  float64 `json:"Score,omitempty"`
}

//ReturnTransition ...
type ReturnTransition struct {
	FromSong    string       `json:"FromSong,omitempty"`
	Transitions []Transition `json:"Transition,omitempty"`
}
