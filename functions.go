package main

import "database/sql"

func convertKey(key, mode int64) (keyString string) {

	var keyS string
	var modeS string

	switch key {
	case 1:
		keyS = "C"
		break
	case 2:
		keyS = "G"
		break
	case 3:
		keyS = "D"
		break
	case 4:
		keyS = "A"
		break
	case 5:
		keyS = "E"
		break
	case 6:
		keyS = "B"
		break
	case 7:
		keyS = "G𝄬/F#"
		break
	case 8:
		keyS = "D𝄬"
		break
	case 9:
		keyS = "A𝄬"
		break
	case 10:
		keyS = "E𝄬"
		break
	case 11:
		keyS = "B𝄬"
		break
	case 12:
		keyS = "F"
		break
	default:
		keyS = "/"
	}

	switch mode {
	case 0:
		modeS = "Dur"
		break
	case 1:
		modeS = "Major"
		break
	default:
		keyS = "/"
	}

	keyString = keyS + " " + modeS

	return
}

func getSongData(songID string) (songObject SongDetailed) {

	db := initDB()
	defer db.Close()
	rows, err := db.Query("select songs.id as id, songs.title as title, artists.name as artist, spotify_songs.tempo as bpm, spotify_songs.key as key, "+
		"spotify_songs.mode as mode, spotify_songs.energy as energy, spotify_songs.instrumentalness as instrumental, spotify_songs.valence, "+
		"spotify_songs.danceability as danceability, spotify_songs.loudness as loudness, spotify_songs.time_signature as timeSignature, "+
		"spotify_songs.duration_ms as duration "+
		"from songs "+
		"join artists on artists.id = songs.artist_id "+
		"join spotify_songs on spotify_songs.song_id = songs.id "+
		"where songs.id = $1 "+
		"group by "+
		"songs.id, songs.title, artists.name, spotify_songs.tempo, spotify_songs.key, spotify_songs.mode, spotify_songs.energy, spotify_songs.instrumentalness, "+
		"spotify_songs.valence, spotify_songs.danceability, spotify_songs.loudness, spotify_songs.time_signature, spotify_songs.duration_ms", songID)
	checkErr(err, "4: Query error!")

	//id, title, artist, previw, cover, numberofsets, bpm, key, rep, energy, instrum, dance, loud, val, timeS, Genre, Dur
	for rows.Next() {
		var id sql.NullString
		var title sql.NullString
		var artist sql.NullString
		var bpm sql.NullFloat64
		var key sql.NullInt64
		var mode sql.NullInt64
		var energy sql.NullFloat64
		var instrumental sql.NullFloat64
		var valence sql.NullFloat64
		var danceability sql.NullFloat64
		var loudness sql.NullFloat64
		var timeSignature sql.NullFloat64
		var duration sql.NullFloat64
		err = rows.Scan(&id, &title, &artist, &bpm, &key, &mode, &energy, &instrumental, &valence, &danceability, &loudness, &timeSignature, &duration)
		checkErr(err, "Corrupt data format!")

		keyString := convertKey(key.Int64, mode.Int64)

		songObject = SongDetailed{
			ID:            id.String,
			Title:         title.String,
			Artist:        artist.String,
			BPM:           bpm.Float64,
			Key:           keyString,
			Reputation:    0,
			Energy:        energy.Float64,
			Instrumental:  instrumental.Float64,
			Danceability:  danceability.Float64,
			Loudness:      loudness.Float64,
			Valence:       valence.Float64,
			TimeSignature: timeSignature.Float64,
			Duration:      duration.Float64,
			Genre:         "",
			PreviewURL:    "",
			CoverURL:      ""}
	}
	return

}

func getTransitionData(fromSong SongDetailed) (transitions []Transition) {

	if fromSong.ID != "" {
		db := initDB()
		defer db.Close()
		rows, err := db.Query("select songs.id as id, songs.title as title, artists.name as artist, spotify_songs.tempo as bpm, spotify_songs.key as key, "+
			"spotify_songs.mode as mode, spotify_songs.energy as energy, spotify_songs.instrumentalness as instrumental, spotify_songs.valence, "+
			"spotify_songs.danceability as danceability, spotify_songs.loudness as loudness, spotify_songs.time_signature as timeSignature, "+
			"spotify_songs.duration_ms as duration, count(songs.id) as occasions "+
			"from songs "+
			"join artists on artists.id = songs.artist_id "+
			"join spotify_songs on spotify_songs.song_id = songs.id "+
			"join transitions on transitions.song_to = songs.id "+
			"where transitions.song_from = $1 "+
			"group by "+
			"songs.id, songs.title, artists.name, spotify_songs.tempo, spotify_songs.key, spotify_songs.mode, spotify_songs.energy, spotify_songs.instrumentalness, "+
			"spotify_songs.valence, spotify_songs.danceability, spotify_songs.loudness, spotify_songs.time_signature, spotify_songs.duration_ms", fromSong.ID)
		checkErr(err, "4: Query error!")

		//id, title, artist, previw, cover, numberofsets, bpm, key, rep, energy, instrum, dance, loud, val, timeS, Genre, Dur
		for rows.Next() {
			var id sql.NullString
			var title sql.NullString
			var artist sql.NullString
			var bpm sql.NullFloat64
			var key sql.NullInt64
			var mode sql.NullInt64
			var energy sql.NullFloat64
			var instrumental sql.NullFloat64
			var valence sql.NullFloat64
			var danceability sql.NullFloat64
			var loudness sql.NullFloat64
			var timeSignature sql.NullFloat64
			var duration sql.NullFloat64
			var occasions sql.NullInt64
			err = rows.Scan(&id, &title, &artist, &bpm, &key, &mode, &energy, &instrumental, &valence, &danceability, &loudness, &timeSignature, &duration, &occasions)
			checkErr(err, "Corrupt data format!")

			keyString := convertKey(key.Int64, mode.Int64)

			toSong := SongDetailed{
				ID:            id.String,
				Title:         title.String,
				Artist:        artist.String,
				BPM:           bpm.Float64,
				Key:           keyString,
				Reputation:    0,
				Energy:        energy.Float64,
				Instrumental:  instrumental.Float64,
				Danceability:  danceability.Float64,
				Loudness:      loudness.Float64,
				Valence:       valence.Float64,
				TimeSignature: timeSignature.Float64,
				Duration:      duration.Float64,
				Genre:         "",
				PreviewURL:    "",
				CoverURL:      ""}

			transitions = append(transitions, Transition{FromSong: fromSong, ToSong: toSong, Occasions: occasions.Int64})
		}
	}
	return

}
