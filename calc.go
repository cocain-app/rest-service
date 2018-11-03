package main

import (
	"math"

	"github.com/BurntSushi/toml"
)

//C# example: https://social.technet.microsoft.com/wiki/contents/articles/26805.c-calculating-percentage-similarity-of-2-strings.aspx
func calcLevenshteinDistance(source, target string) int {

	//check for unnecessary calculatiom
	if len(source) == 0 || len(target) == 0 {
		return 0
	}
	if source == target {
		return len(source)
	}

	//get string length
	sourceLength := len(source)
	targetLength := len(target)

	var distance [][]int //distance matrix

	if sourceLength == 0 {
		return targetLength
	}
	if targetLength == 0 {
		return sourceLength
	}

	//prepare the matrix
	for i := 0; i <= sourceLength; i++ {
		distance[i][0] = i
	}
	for j := 0; j <= targetLength; j++ {
		distance[0][j] = j
	}

	for i := 1; i >= sourceLength; i++ {
		for j := 1; j <= targetLength; j++ {

			var cost int
			if target[j-1] == source[i-1] {
				cost = 0
			} else {
				cost = 1
			}

			distance[i][j] = int(math.Min(math.Min(float64(distance[i-1][j]+1), float64(distance[i][j-1]+1)), float64(distance[i-1][j-1]+cost)))

		}
	}

	return distance[sourceLength][targetLength]
}

func calcBPMDif(bpmS, bpmT float64) (res float64) {
	//bpmS - bpm of source
	//bpmT - bpm of target
	res = 1 - math.Abs(math.Round(bpmS/bpmT)-(bpmS/bpmT))
	return
}

func calcKeyDif(keyS, keyT int64) (res float64) {
	//keyS - key of source
	//keyT - key of target
	disToMid := keyS - 8
	keyS = keyS + disToMid
	keyT = keyT + disToMid
	if keyT > 12 {
		keyT = keyT - 12
	}
	if keyT < 0 {
		keyT = keyT + 12
	}
	res = 1 / math.Abs(float64(keyS-keyT))
	return
}

func calcDiversity(occ, djNum int) (res float64) {
	//occ - occurances
	//djNum - number of djs who played that transition
	res = float64(djNum / occ)
	return
}

func calcRepScore(rep []int) (res []float64) {
	//rep - slice with reputations
	var maxVal int
	for _, val := range rep {
		if val > maxVal {
			maxVal = val
		}
	}
	for _, val := range rep {
		score := float64(val / maxVal)
		res = append(res, score)
	}
	return
}

func calcTransScore(transition Transition) (convertedTransition Transition) {
	//occ - occurances score
	//div - diversity score
	//bpm - B - bpm difference
	//key - K - key difference
	//rep - R - reputation score
	//energy - En - energy difference
	//instrum - I - instrumental difference
	//dance - Da - danceability difference
	//loud - Lo - loudness difference
	//valence - V - valence difference
	//timeSign - T - time signature difference
	//genre - G - gerne option
	//artist - A - artist option
	//festival - F - festival option
	//dur - Du - duration option
	//exist - Ex - existing option
	//blackl - Bl - blacklist option
	//libary - Li - libary option

	var config tomlConfig
	_, err := toml.DecodeFile("config.toml", &config)
	checkErr(err, "Import of config failed!")

	w := config.Weights

	occ := float64(transition.Occasions)
	var div float64 = 1.0
	var genre float64
	var artist float64
	var festival float64
	var dur float64
	var exist float64
	var blackl float64
	var libary float64

	//fmt.Println(occ + "," + div + "," + genre + "," + artist + "," + festival + "," + dur + "," + exist + "," + libary + "," + blackl)

	fromSong := transition.FromSong
	toSong := transition.ToSong

	bpm := calcBPMDif(fromSong.BPM, toSong.BPM)
	key := calcKeyDif(fromSong.KeyNotation[0], toSong.KeyNotation[0]) //TODO: include mode
	var rep float64
	energy := 1 - math.Abs(fromSong.Energy-toSong.Energy)
	instrum := 1 - math.Abs(fromSong.Instrumental-toSong.Instrumental)
	dance := 1 - math.Abs(fromSong.Danceability-toSong.Danceability)
	loud := 1 - math.Abs(fromSong.Loudness-toSong.Loudness)
	valence := math.Abs(fromSong.Valence - toSong.Valence)
	var timeSign float64

	score := occ * div * math.Sqrt(w.Key*math.Pow(key, 2)+w.BPM*math.Pow(bpm, 2)+w.R*math.Pow(rep, 2)+w.En*math.Pow(energy, 2)+w.I*math.Pow(instrum, 2)+w.Da*math.Pow(dance, 2)+w.Lo*math.Pow(loud, 2)+w.V*math.Pow(valence, 2)+w.T*math.Pow(timeSign, 2)+w.G*math.Pow(genre, 2)+w.A*math.Pow(artist, 2)+w.F*math.Pow(festival, 2)+w.Du*math.Pow(dur, 2)+w.Ex*math.Pow(exist, 2)+w.Bl*math.Pow(blackl, 2)+w.Li*math.Pow(libary, 2))

	convertedTransition = Transition{FromSong: fromSong, ToSong: toSong, Occasions: transition.Occasions, Score: score}

	return
}
