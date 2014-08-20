package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestWord("Hello my name is Martin"))
	fmt.Println(longestWord("San Francisco is sunny in the summer. Sort of..."))
}

func longestWord(sen string) string {
	senSlice := strings.Split(sen, " ")
	longestWord := senSlice[0]
	for i := range senSlice {
		if len(senSlice[i]) > len(longestWord) {
			longestWord = senSlice[i]
		}
	}
	return longestWord
}
