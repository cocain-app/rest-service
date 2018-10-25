package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

func getSuggestions(w http.ResponseWriter, r *http.Request) {
	songID := r.Header.Get("songID")

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
	songID := ""

	db := initDB()
	defer db.Close()
	rows, err := db.Query("SELECT id FROM public.songs WHERE title ILIKE $1", songTitle)
	checkErr(err, "2: Query error!")

	for rows.Next() {
		var id sql.NullString
		err = rows.Scan(&id)
		checkErr(err, "Corrupt data format!")

		songID = id.String //Catch multiple ids
	}

	json.NewEncoder(w).Encode(songID)
}

/*
func getTransitions(songID string, db *sql.DB) (transitions []string, err error) {
	rows, err := db.Query("SELECT song_to FROM public.transitions WHERE song_from=$1", songID)
	checkErr(err, "3: Query error!")

	for rows.Next() {
		var song_to sql.NullString
		err = rows.Scan(&song_to)
		checkErr(err, "Corrupt data format!")

		transitions = append(transitions, song_to.String) //Catch multiple ids
	}

	if len(transitions) <= 0 {
		err = errors.New("sql error: no query results")
	}
	return
}

func getArtist(artistID string, db *sql.DB) (artistName string, err error) {
	rows, err := db.Query("SELECT name FROM public.artists WHERE id=$1", artistID)
	checkErr(err, "4: Query error!")

	for rows.Next() {
		var name sql.NullString
		err = rows.Scan(&name)
		checkErr(err, "Corrupt data format!")

		artistName = name.String //Catch multiple ids
	}

	if artistID == "" {
		err = errors.New("sql error: no query results")
	}
	return
}
*/
