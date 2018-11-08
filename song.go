package main

import (
	_ "github.com/lib/pq"
)

//SearchSong ...
type SearchSong struct {
	Song    Song  `json:"Song,omitempty"`
	LenDiff int64 `json:"LenDiff,omitempty"`
}

//Song ...
type Song struct {
	ID            string  `json:"ID,omitempty"`
	Title         string  `json:"Title,omitempty"`
	Artist        string  `json:"Artist,omitempty"`
	BPM           float64 `json:"BPM,omitempty"`
	Key           string  `json:"Key,omitempty"`
	Duration      int64   `json:"Duration,omitempty"`
	PreviewURL    string  `json:"PreviewURL,omitempty"`
	ImageURL      string  `json:"ImageURL,omitempty"`
	ImageURLSmall string  `json:"ImageURLSmall,omitempty"`
}

//SongDetailed ...
type SongDetailed struct {
	ID            string   `json:"ID,omitempty"`
	Title         string   `json:"Title,omitempty"`
	Artist        string   `json:"Artist,omitempty"`
	PreviewURL    string   `json:"PreviewURL,omitempty"`
	ImageURL      string   `json:"ImageURL,omitempty"`
	ImageURLSmall string   `json:"ImageURLSmall,omitempty"`
	BPM           float64  `json:"BPM,omitempty"`
	KeyNotation   [2]int64 `json:"KeyNotation,omitempty"`
	Key           string   `json:"Key,omitempty"`
	Reputation    int64    `json:"Reputation,omitempty"`
	Energy        float64  `json:"Energy,omitempty"`
	Instrumental  float64  `json:"Instrumental,omitempty"`
	Danceability  float64  `json:"Danceability,omitempty"`
	Loudness      float64  `json:"Loudness,omitempty"`
	Valence       float64  `json:"Valence,omitempty"`
	TimeSignature float64  `json:"TimeSignature,omitempty"`
	Genre         string   `json:"Genre,omitempty"`
	Duration      int64    `json:"Duration,omitempty"`
}
