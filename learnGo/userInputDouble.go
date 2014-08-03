package main

import "fmt"

func main() {
	fmt.Println("Please enter a number: ")
	var num int
	fmt.Scanf("%d", &num)
	fmt.Println(num * 2)
}
