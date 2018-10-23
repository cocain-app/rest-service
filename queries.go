package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func getSuggestions(w http.ResponseWriter, r *http.Request) {

	songTitle := "The Way"
	db := initDB()
	defer db.Close()
	songID := getSongID(songTitle, db)

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
		json.NewEncoder(w).Encode(suggestions)
	*/
}

func getSongID(songTitle string, db *sql.DB) (songID string) {
	rows, err := db.Query("SELECT id FROM public.songs WHERE title=$1", songTitle)
	checkErr(err, "Query error!")

	songID = ""

	for rows.Next() {
		var id sql.NullString
		err = rows.Scan(&id)
		checkErr(err, "Corrupt data format!")

		songID = id.String //Catch multiple ids
	}

	return
}
