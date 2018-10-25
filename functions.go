package main

//C# example: https://social.technet.microsoft.com/wiki/contents/articles/26805.c-calculating-percentage-similarity-of-2-strings.aspx
func calcLevenshteinDistance(source, target string) (distance int) {

	if len(source) == 0 || len(target) == 0 {
		return 0
	}
	if source == target {
		return len(source)
	}

	sourceLength := len(source)
	targetLength := len(target)

	if sourceLength == 0 {
		return targetLength
	}
	if targetLength == 0 {
		return sourceLength
	}

	return
}
