package main

import (
	_ "github.com/lib/pq"
)

//SongSuggestion object for returning songlist
type SongSuggestion struct {
	ID           string  `json:"ID,omitempty"`
	Title        string  `json:"Title,omitempty"`
	Artist       string  `json:"Artist,omitempty"`
	NumberOfSets int64   `json:"NumberOfSets,omitempty"`
	LenDiff      float64 `json:"LenDiff,omitempty"`
}

//bpm, key, energy, previewURL, coverURL, durration
type Song struct {
	ID         string  `json:"ID,omitempty"`
	Title      string  `json:"Title,omitempty"`
	Artist     string  `json:"Artist,omitempty"`
	BPM        float64 `json:"BPM,omitempty"`
	Key        string  `json:"Key,omitempty"`
	Energy     float64 `json:"Energy,omitempty"`
	Duration   float64 `json:"Duration,omitempty"`
	PreviewURL string  `json:"PreviewURL,omitempty"`
	CoverURL   string  `json:"CoverURL,omitempty"`
}

//occ, div, bpm, key, rep, energy, instrum, dance, loud, valence, timeSign, genre, artist, festival, dur, exist, blackl, libary
type SongDetailed struct {
	ID            string  `json:"ID,omitempty"`
	Title         string  `json:"Title,omitempty"`
	Artist        string  `json:"Artist,omitempty"`
	PreviewURL    string  `json:"PreviewURL,omitempty"`
	CoverURL      string  `json:"CoverURL,omitempty"`
	BPM           float64 `json:"BPM,omitempty"`
	Key           string  `json:"Key,omitempty"`
	Reputation    int64   `json:"Reputation,omitempty"`
	Energy        float64 `json:"Energy,omitempty"`
	Instrumental  float64 `json:"Instrumental,omitempty"`
	Danceability  float64 `json:"Danceability,omitempty"`
	Loudness      float64 `json:"Loudness,omitempty"`
	Valence       float64 `json:"Valence,omitempty"`
	TimeSignature float64 `json:"TimeSignature,omitempty"`
	Genre         string  `json:"Genre,omitempty"`
	Duration      float64 `json:"Duration,omitempty"`
}
