package main

import (
	"fmt"
	"io/ioutil"
	"log"
	s "strings"
)

func main() {
	file, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	mys := string(file)
	wordArray := s.Fields(mys)
	newSlice := make([]string, len(wordArray)+1)
	copy(newSlice, wordArray)
	for key, value := range wordArray {
		if value == "word" {
			newSlice = Insert(newSlice, key, "another")
		}
	}
	fmt.Println(newSlice)

	ioutil.WriteFile("appendedFile.txt", []byte(s.Join(newSlice, " ")), 222)
}

func Insert(slice []string, index int, value string) []string {
	slice = slice[0:len(slice)]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}
