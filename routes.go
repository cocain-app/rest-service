package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

/*----------------------------------------------------
Transitions - /songs/transitions
----------------------------------------------------*/
func getTransitions(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]
	fmt.Println(songID)

	var plainTransitions []TransitionDetailed
	songData := getSongData(songID)
	plainTransitions = getTransitionData(songData)

	for i, transition := range plainTransitions {
		if transition.FromSong.BPM != 0 && transition.ToSong.BPM != 0 {
			plainTransitions[i] = calcTransScore(transition)
		} else {
			plainTransitions[i].Score = 0
		}
	}

	var transitions []Transition

	for _, t := range plainTransitions {
		toSong := simSong(t.ToSong)
		transitions = append(transitions, Transition{ToSong: toSong, Score: t.Score})
	}

	sort.Slice(transitions, func(i, j int) bool {
		return transitions[i].Score > transitions[j].Score
	})

	rTransitions := ReturnTransition{FromSong: songID, Transitions: transitions}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rTransitions)
}

/*----------------------------------------------------
Song Search - /search
----------------------------------------------------*/
func getSongs(w http.ResponseWriter, r *http.Request) {
	query := r.Header.Get("query")
	fmt.Println(query)
	fromBpm := r.Header.Get("fromBpm")
	toBpm := r.Header.Get("toBpm")
	key := r.Header.Get("key")
	mode := r.Header.Get("mode")

	songs := searchQuery(query, fromBpm, toBpm, key, mode)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

/*----------------------------------------------------
Song Details - /song/get
----------------------------------------------------*/
func getSongDetails(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]
	fmt.Println(songID)

	db := initDB()
	defer db.Close()
	rows, err := db.Query("select songs.title as title, artists.name as artist, spotify_songs.tempo as bpm, spotify_songs.key as key, "+
		"spotify_songs.mode as mode, spotify_songs.duration_ms as duration, spotify_songs.preview_url as previewURL, spotify_songs.image_url_large as imageURL, spotify_songs.image_url_small as imageURLSmall "+
		"from songs join artists on artists.id = songs.artist_id "+
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
		var duration sql.NullInt64
		var previewURL sql.NullString
		var imageURL sql.NullString
		var imageURLSmall sql.NullString
		err = rows.Scan(&title, &artist, &bpm, &key, &mode, &duration, &previewURL, &imageURL, &imageURLSmall)
		checkErr(err, "Corrupt data format!")

		keyString := convertKey(key.Int64, mode.Int64)

		song = Song{
			ID:            songID,
			Title:         title.String,
			Artist:        artist.String,
			BPM:           round(bpm.Float64, 0.5),
			Key:           keyString,
			Duration:      duration.Int64,
			PreviewURL:    previewURL.String,
			ImageURL:      imageURL.String,
			ImageURLSmall: imageURLSmall.String}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

/*----------------------------------------------------
DEV: All Song Details - /songs/get/.../all
----------------------------------------------------*/
func getAllSongDetails(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]
	fmt.Println(songID)

	song := getSongData(songID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

func isOnline(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Pong")
}
