package main

func MaximumTime(initialTime string) string {

	// "01:23"
	// "23:59"
	runeTime := []rune(initialTime)
	for GetIndex(runeTime, '?') != -1 {
		switch GetIndex(runeTime, '?') {
		case 4:
			runeTime[4] = '9'
		case 3:
			runeTime[3] = '5'
		case 1:
			if runeTime[0] == '1' || runeTime[0] == '0' {
				runeTime[1] = '9'
			} else {
				runeTime[1] = '3'
			}
		case 0:
			if runeTime[1] != '0' &&
				runeTime[1] != '1' &&
				runeTime[1] != '2' &&
				runeTime[1] != '3' &&
				runeTime[1] != '?' {
				runeTime[0] = '1'
			} else {
				runeTime[0] = '2'
			}
		}
	}
	return string(runeTime)
}

func GetIndex(s []rune, char rune) int {
	for i := range s {
		if s[i] == char {
			return i
		}
	}
	return -1
}

// "01234"
// "22:??" -> 22:59
// "2?:??" -> 23:59
// "2?:5?" -?
