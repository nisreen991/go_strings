package impl

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	const CHAR int = 256
	var count [CHAR]int

	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for i := 0; i < CHAR; i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
