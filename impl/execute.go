package impl

import "fmt"

func Execute() {
	s := "Hello World "
	s1 := "-India-is-a-great-country-"
	substr := "World"
	s2 := "hello world "
	s3 := "abbb"
	fmt.Println(containsSubstring(s, substr))
	fmt.Println(upperCase(s))
	fmt.Println(lowerCase(s))
	fmt.Println(index(s, substr))
	fmt.Println(length(s))
	fmt.Println(split(s1))
	fmt.Print(repeat(s))
	fmt.Println(replace(s))
	fmt.Println(join(split(s1)))
	fmt.Println(hasPrefix(s))
	fmt.Println(hasSuffix(s))
	fmt.Println(countChar(s))
	printChar(s)
	printAscii(s)
	fmt.Println(trimSpace(s))
	fmt.Println(trimChar(s1))
	fmt.Println(trimLeft(s1))
	fmt.Println(trimRight(s1))
	fmt.Println(compareString(s, s1))
	fmt.Println(containsChar(s1))
	fmt.Println(containsUniCode(s))
	fmt.Println(matchStrings(s, s2))
	fmt.Println(indexByte(s))
	fmt.Println(lastOccurenceIndex(s, substr))
	fmt.Println(lastOccurenceByte(s))
	fmt.Println(replaceAll(s))
	fmt.Println(titleCase(s1))
	fmt.Println(palindromeCheck(s3))
	fmt.Print(reverseString(s3))
}
