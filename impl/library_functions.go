package impl

import (
	"fmt"
	"strings"
)

func containsSubstring(s string, substr string) bool {
	return strings.Contains(s, substr)
}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func lowerCase(s string) string {
	return strings.ToLower(s)
}

func index(s string, substr string) int {
	return strings.Index(s, substr)
}

func length(s string) int {
	return len(s)
}

func split(s string) []string {
	return strings.Split(s, "-")
}

func repeat(s string) string {
	return strings.Repeat(s, 4)
}

func replace(s string) string {
	return strings.Replace(s, "W", "w", 4) // learn more about this
}

func join(s []string) string {
	return strings.Join(s, "-")
}

func hasPrefix(s string) bool {
	return strings.HasPrefix(s, "He")
}

func hasSuffix(s string) bool {
	return strings.HasSuffix(s, "rld")
}

func countChar(s string) int {
	return strings.Count(s, "o")
}

func printChar(s string) {
	for _, ch := range s {
		fmt.Println(string(ch))
	}
}

func printAscii(s string) {
	for _, ch := range s {
		fmt.Println(string(ch), ch)
	}
}

func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func trimChar(s string) string {
	return strings.Trim(s, "-")
}

func trimLeft(s string) string {
	return strings.TrimLeft(s, "-")
}

func trimRight(s string) string {
	return strings.TrimRight(s, "-")
}

func compareString(s1, s2 string) int {
	return strings.Compare(s1, s2)
}

func containsChar(s string) bool {
	return strings.ContainsAny(s, "r")
}

func containsUniCode(s string) bool {
	return strings.ContainsRune(s, 68)
}

func matchStrings(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

func indexByte(s string) int {
	return strings.IndexByte(s, 'o')
}

func lastOccurenceIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

func lastOccurenceByte(s string) int {
	return strings.LastIndexByte(s, 'r')
}

func replaceAll(s string) string {
	return strings.ReplaceAll(s, "World", "Kuwait")
}

func titleCase(s string) string {
	return strings.Title(s)
}
