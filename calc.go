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

func calcKeyDif(keyS, keyT int) (res float64) {
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

func calcTransScore(occ, div, bpm, key, rep, energy, instrum, dance, loud, valence, timeSign, genre, artist, festival, dur, exist, blackl, libary float64) (score float64) {
	//occ - occurances score
	//div - diversity score
	//bpm - bpm difference
	//key - key difference
	//rep - reputation score
	//energy - energy difference
	//instrum - instrumental difference
	//dance - danceability difference
	//loud - loudness difference
	//valence - valence difference
	//timeSign - time signature difference
	//genre - gerne option
	//artist - artist option
	//festival - festival option
	//dur - duration option
	//exist - existing option
	//blackl - blacklist option
	//libary - libary option

	var w Weights
	_, err := toml.Decode("config.toml", &w)
	checkErr(err, "Import of config failed!")

	score = occ * div * math.Sqrt(w.K*math.Pow(key, 2)+w.B*math.Pow(bpm, 2)+w.R*math.Pow(rep, 2)+w.En*math.Pow(energy, 2)+w.I*math.Pow(instrum, 2)+w.Da*math.Pow(instrum, 2)+w.Lo*math.Pow(loud, 2)+w.V*math.Pow(valence, 2)+w.T*math.Pow(timeSign, 2)+w.G*math.Pow(genre, 2)+w.A*math.Pow(artist, 2)+w.F*math.Pow(festival, 2)+w.Du*math.Pow(dur, 2)+w.Ex*math.Pow(exist, 2)+w.Bl*math.Pow(blackl, 2)+w.Li*math.Pow(libary, 2))

	return

}
