package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func getSuggestions(w http.ResponseWriter, r *http.Request) {

	songTitle := "Losing It"
	db := initDB()
	defer db.Close()
	songID, err := getSongID(songTitle, db) //err = no results
	checkErr(err, "No query results")
	transitions, err := getTransitions(songID, db) //err= no results
	checkErr(err, "No query results")

	fmt.Println(transitions)

	json.NewEncoder(w).Encode(transitions)

	/*
		rows, err := db.Query("SELECT id, title, artist_id, release_date, duration, spotify_uri, timestamp_added, timestamp_modified FROM public.songs WHERE title=$1", songTitle)
			checkErr(err, "Query error!")

			var suggestions []Suggestion


			for rows.Next() {
				var id sql.NullString
				var title sql.NullString
				var artist_id sql.NullString
				var release_date sql.NullString
				var duration sql.NullInt64
				var spotify_uri sql.NullString
				var timestamp_added sql.NullString
				var timestamp_modified sql.NullString
				err = rows.Scan(&id, &title, &artist_id, &release_date, &duration, &spotify_uri, &timestamp_added, &timestamp_modified)
				checkErr(err, "Corrupt data format!")
				//fmt.Println(" id | title | artist_id | release_date | duration | spotify_uri | timestampt_added | timestamp_modified ")
				fmt.Println(id, title, artist_id, release_date, duration, spotify_uri, timestamp_added, timestamp_modified)

				suggestions = append(suggestions, Suggestion{ID: id.String, Title: title.String})
			}

			fmt.Println(suggestions)

	*/
}

func getSongID(songTitle string, db *sql.DB) (songID string, err error) {
	rows, err := db.Query("SELECT id FROM public.songs WHERE title=$1", songTitle)
	checkErr(err, "2: Query error!")

	for rows.Next() {
		var id sql.NullString
		err = rows.Scan(&id)
		checkErr(err, "Corrupt data format!")

		songID = id.String //Catch multiple ids
	}

	if songID == "" {
		err = errors.New("sql error: no query results")
	}
	return
}

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
