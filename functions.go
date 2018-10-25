package main

import (
	"math"
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
