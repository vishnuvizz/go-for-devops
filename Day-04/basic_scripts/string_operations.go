package main

import (
	"fmt"
	"strings"
)

func main() {
	// String Concatenation
	str1 := "Hello"
	str2 := "World"
	concatenated := str1 + " " + str2
	fmt.Println("Concatenated String:", concatenated)

	// string concat by Vishnu
	str3 := "Hi"
	str4 := "there"
	Concat := str3 + " " + str4
	fmt.Println("Concatenated String:", Concat)

	// Splitting a string
	sentence := "Go is an open source programming language."
	words := strings.Split(sentence, " ")
	fmt.Println("Words in sentence:", words)

	//Splitting string by Vishnu
	sentence1 := "Hi there how is everything going"
	words1 := strings.Split(sentence1, " ")
	fmt.Println("Words in sentence:", words1)

	// Replacing a substring
	replaced := strings.Replace(sentence, "Go", "Golang", 1)
	fmt.Println("Replaced String:", replaced)

	//replacing a substring
	replaced1 := strings.Replace(sentence1, "everything", "all", 1)
	fmt.Println("replacing string by Vishnu:", replaced1)

	// Upper and Lower case
	upper := strings.ToUpper(sentence)
	lower := strings.ToLower(sentence)
	fmt.Println("Uppercase:", upper)
	fmt.Println("Lowercase:", lower)

	// Trimming
	spaceyString := "   trim me   "
	trimmed := strings.TrimSpace(spaceyString)
	fmt.Println("Trimmed String:", trimmed)

	// Substring - using slicing
	// Note: Go does not have a built-in 'substring' method, but we can achieve this with slicing
	start := 3
	end := 11
	if end <= len(sentence) && start < end {
		substring := sentence[start:end]
		fmt.Println("Substring:", substring)
	} else {
		fmt.Println("Invalid range for substring")
	}

	// Checking if a string contains a substring
	contains := strings.Contains(sentence, "source")
	fmt.Println("Contains 'source':", contains)

	// Finding a substring's index
	index := strings.Index(sentence, "source")
	fmt.Println("Index of 'source':", index)
}
