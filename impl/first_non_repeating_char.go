package impl

func firstNonRepeatingChar(s string) string {
	charMap := make(map[string]int)

	for _, ch := range s {
		charMap[string(ch)]++
	}
	for key, value := range charMap {
		if value == 1 {
			return key
		}
	}
	return "$"
}
