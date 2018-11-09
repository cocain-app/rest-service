package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func searchQuery(queryS, bpm1, bpm2, key, mode string) (songs []SearchSong) {

	query := strings.Split(queryS, " ")

	db := initDB()
	defer db.Close()

	var searchArray string

	var sb strings.Builder

	for i, queryPart := range query {
		if i > 0 {
			sb.WriteString("'%")
		} else {
			sb.WriteString("'")
		}
		sb.WriteString(queryPart)
		sb.WriteString("%'")
		if i < len(query)-1 {
			sb.WriteString(",")
		}
	}

	searchArray = sb.String()
	fmt.Println(searchArray)
	sb.Reset()
	options := ""

	if len(bpm1) != 0 {
		sb.WriteString("and spotify_songs.tempo between ")
		bpm1Value, err := strconv.ParseFloat(bpm1, 64)
		checkErr(err, "Invalid data type for fromBpm")
		if len(bpm2) != 0 {
			bpm2Value, err := strconv.ParseFloat(bpm2, 64)
			checkErr(err, "Invalid data type for fromBpm")
			sb.WriteString(fmt.Sprintf("%f", bpm1Value-0.25))
			sb.WriteString(" and ")
			sb.WriteString(fmt.Sprintf("%f", bpm2Value+0.25))
		} else {
			sb.WriteString(fmt.Sprintf("%f", bpm1Value-0.5))
			sb.WriteString(" and ")
			sb.WriteString(fmt.Sprintf("%f", bpm1Value+0.5))
		}
		sb.WriteString(" ")
	}

	if len(key) != 0 {
		if len(query) > 0 || len(bpm1) > 0 {
			sb.WriteString("and ")
		}
		sb.WriteString("key = ")
		sb.WriteString(key)
		sb.WriteString(" ")
	}

	if len(mode) != 0 {
		if len(query) > 0 || len(key) > 0 {
			sb.WriteString("and ")
		}
		sb.WriteString("mode = ")
		sb.WriteString(mode)
		sb.WriteString(" ")
	}

	options = sb.String()

	fmt.Println(options)

	sqlQuery := fmt.Sprintf("select songs.id as id, songs.title as title, artists.name as artist, spotify_songs.tempo as bpm, "+
		"spotify_songs.key as key, spotify_songs.mode as mode, spotify_songs.duration_ms as duration, "+
		"spotify_songs.preview_url as previewURL, spotify_songs.image_url_large as imageURL, spotify_songs.image_url_small as imageURLSmall, "+
		"levenshtein('%s',title) as diff, count(songs.id)-1 as score "+
		"from songs join artists on artists.id = songs.artist_id join spotify_songs on spotify_songs.song_id = songs.id "+
		"join transitions on transitions.song_from = songs.id "+
		"where (songs.title || ' ' || artists.name ilike ALL(Array[%s]) or artists.name || ' ' || songs.title ilike ALL(Array[%s])) %s "+
		"group by songs.id, artists.id, spotify_songs.tempo, spotify_songs.key, "+
		"spotify_songs.mode, spotify_songs.preview_url, spotify_songs.image_url_large, spotify_songs.image_url_small, "+
		"spotify_songs.duration_ms "+
		"order by score desc, diff asc limit 20", query, searchArray, searchArray, options)

	rows, err := db.Query(sqlQuery)

	checkErr(err, "2: Query error!")

	for rows.Next() {
		var id sql.NullString
		var title sql.NullString
		var artist sql.NullString
		var bpm sql.NullFloat64
		var key sql.NullInt64
		var mode sql.NullInt64
		var duration sql.NullInt64
		var previewurl sql.NullString
		var imageurl sql.NullString
		var imageurlsmall sql.NullString
		var diff sql.NullInt64
		var score sql.NullInt64
		err = rows.Scan(&id, &title, &artist, &bpm, &key, &mode, &duration, &previewurl, &imageurl, &imageurlsmall, &diff, &score)
		checkErr(err, "Corrupt data format!")

		keyString := convertKey(key.Int64, mode.Int64)

		songs = append(songs, SearchSong{
			Song: Song{
				ID:            id.String,
				Title:         title.String,
				Artist:        artist.String,
				BPM:           round(bpm.Float64, 0.5),
				Key:           keyString,
				Duration:      duration.Int64,
				PreviewURL:    previewurl.String,
				ImageURL:      imageurl.String,
				ImageURLSmall: imageurlsmall.String},
			LenDiff:  diff.Int64,
			TransNum: score.Int64})
	}
	return
}
