package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func getTransitions(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]

	var transitions []Transition
	songData := getSongData(songID)
	transitions = getTransitionData(songData)
	json.NewEncoder(w).Encode(transitions)
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	songTitle := r.Header.Get("songTitle")
	queryTitle := "%" + songTitle + "%"
	var ids []SongSuggestion
	titleLength := len(songTitle)

	db := initDB()
	defer db.Close()
	rows, err := db.Query("SELECT songs.id AS id, songs.title AS title, artists.name AS artistName, count(songs.id) AS score FROM songs JOIN artists ON artists.id = songs.artist_id "+
		"WHERE title ILIKE $1 GROUP BY songs.id, artists.id ORDER BY score desc LIMIT 10 ", queryTitle)
	checkErr(err, "2: Query error!")

	for rows.Next() {
		var id sql.NullString
		var title sql.NullString
		var artistName sql.NullString
		var score sql.NullInt64
		err = rows.Scan(&id, &title, &artistName, &score)
		checkErr(err, "Corrupt data format!")

		ids = append(ids, SongSuggestion{ID: id.String, Title: title.String, Artist: artistName.String, NumberOfSets: score.Int64, LenDiff: math.Abs(float64(titleLength - len(title.String)))})
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i].LenDiff < ids[j].LenDiff
	})

	json.NewEncoder(w).Encode(ids)
}

func getSongDetails(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]
	fmt.Println(songID)

	db := initDB()
	defer db.Close()
	rows, err := db.Query("select songs.title as title, artists.name as artist, spotify_songs.tempo as bpm, spotify_songs.key as key, "+
		"spotify_songs.mode as mode, spotify_songs.energy as energy, spotify_songs.duration_ms as duration from songs join artists on artists.id = songs.artist_id "+
		"join spotify_songs on spotify_songs.song_id = songs.id "+
		"where songs.id = $1", songID)
	checkErr(err, "3: Query error!")

	var song Song

	for rows.Next() {
		var title sql.NullString
		var artist sql.NullString
		var bpm sql.NullFloat64
		var key sql.NullInt64
		var mode sql.NullInt64
		var energy sql.NullFloat64
		var duration sql.NullFloat64
		err = rows.Scan(&title, &artist, &bpm, &key, &mode, &energy, &duration)
		checkErr(err, "Corrupt data format!")

		keyString := convertKey(key.Int64, mode.Int64)
		fmt.Println(title.String)

		song = Song{
			ID:         songID,
			Title:      title.String,
			Artist:     artist.String,
			BPM:        bpm.Float64,
			Key:        keyString,
			Energy:     energy.Float64,
			Duration:   duration.Float64,
			PreviewURL: "",
			CoverURL:   ""}
	}

	json.NewEncoder(w).Encode(song)
}

func isOnline(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Pong")
}
