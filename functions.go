package main

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
