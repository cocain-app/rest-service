package main

import (
	_ "github.com/lib/pq"
)

//Suggestion object for returning songlist
type Suggestion struct {
	ID    string `json: "ID,omitempty"`
	Title string `json: "Title,omitempty"`
}
