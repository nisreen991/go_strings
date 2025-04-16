package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Hello World"
	substr := "World"
	fmt.Println(containsSubstring(s, substr))
	fmt.Println(upperCase(s))
	fmt.Println(lowerCase(s))
	fmt.Println(index(s, substr))
	fmt.Println(length(s))
}

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
