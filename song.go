package main

import (
	_ "github.com/lib/pq"
)

//Song object for returning songlist
type Song struct {
	ID           string  `json:"ID,omitempty"`
	Title        string  `json:"Title,omitempty"`
	Artist       string  `json:"Artist,omitempty"`
	NumberOfSets int64   `json:"NumberOfSets,omitempty"`
	LenDiff      float64 `json:"LenDiff,omitempty"`
}
