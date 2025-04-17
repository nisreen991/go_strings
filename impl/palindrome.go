package impl

func palindromeCheck(s string) bool {
	start := 0
	end := len(s) - 1

	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
