package impl

import "strings"

func reverseWordsInString(s string) string {
	words := strings.Fields(s)
	n := len(words)

	start := 0
	end := n - 1

	for start <= end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}

	return strings.Join(words, " ")
}
