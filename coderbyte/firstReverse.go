package main

import "fmt"

func main() {
	firstReverse("My name is Martin")
	firstReverse("Today is Friday!")
}

func firstReverse(s string) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, i-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
}
