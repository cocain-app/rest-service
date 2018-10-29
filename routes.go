package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func getSuggestions(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]

	db := initDB()
	defer db.Close()
	rows, err := db.Query("select songs.id as id, songs.title as title, artists.name as artistName, count(songs.id) as score from songs join artists on artists.id = songs.artist_id "+
		"join transitions on transitions.song_to=songs.id where transitions.song_from=$1 "+
		"group by songs.id, artists.id order by score desc", songID)
	checkErr(err, "Query error!")

	var suggestions []Song

	for rows.Next() {
		var id sql.NullString
		var title sql.NullString
		var artistName sql.NullString
		var score sql.NullInt64
		err = rows.Scan(&id, &title, &artistName, &score)
		checkErr(err, "Corrupt data format!")

		suggestions = append(suggestions, Song{ID: id.String, Title: title.String, Artist: artistName.String, NumberOfSets: score.Int64})
	}

	json.NewEncoder(w).Encode(suggestions)
}

func getSongID(w http.ResponseWriter, r *http.Request) {
	songTitle := r.Header.Get("songTitle")
	queryTitle := "%" + songTitle + "%"
	var ids []Song

	db := initDB()
	defer db.Close()
	rows, err := db.Query("SELECT songs.id AS id, songs.title AS title, artists.name AS artistName, count(songs.id) AS score FROM songs JOIN artists ON artists.id = songs.artist_id "+
		"WHERE title ILIKE $1 ORDER BY score DESC LIMIT 10 ", queryTitle)
	checkErr(err, "2: Query error!")

	for rows.Next() {
		var id sql.NullString
		var title sql.NullString
		var artistName sql.NullString
		var score sql.NullInt64
		err = rows.Scan(&id, &title, &artistName, &score)
		checkErr(err, "Corrupt data format!")

		ids = append(ids, Song{ID: id.String, Title: title.String, Artist: artistName.String, NumberOfSets: score.Int64})
	}

	json.NewEncoder(w).Encode(ids)
}

func getSongData(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]
	//TODO
	/*
		db := initDB()
		defer db.Close()
		rows, err := db.Query("SELECT title FROM public.songs WHERE id=$1", songID)
		checkErr(err, "2: Query error!")

		for rows.Next() {
			var id sql.NullString
			err = rows.Scan(&id)
			checkErr(err, "Corrupt data format!")

			songID = id.String //Catch multiple ids
		}
	*/
	json.NewEncoder(w).Encode(songID)
}
