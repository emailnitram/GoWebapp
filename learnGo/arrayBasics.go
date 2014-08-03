package main

import "fmt"

func main() {
	numbersArray := [5]int{1, 2, 3, 4, 5}
	for i := len(numbersArray) - 1; i >= 0; i-- {
		fmt.Println(numbersArray[i])
	}
	sum := 0
	for _, value := range numbersArray {
		sum += value
	}
	fmt.Println(sum)

	testScores := [6]float64{75.0, 60.0, 99.0, 65.0, 70.0, 88.0}
	// var newSum float64
	newSum := float64(0)
	for _, value := range testScores {
		newSum += value
	}
	fmt.Println(newSum / float64(len(testScores)))
}
