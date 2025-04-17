package impl

import "sort"

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	firstStr := strs[0]
	lastStr := strs[len(strs)-1]

	minlen := min(len(firstStr), length(lastStr))
	var longestPrefix string
	for i := 0; i < minlen; i++ {
		if firstStr[i] == lastStr[i] {
			longestPrefix += string(firstStr[i])
		}
	}

	return longestPrefix
}
