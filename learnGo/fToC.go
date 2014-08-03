package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Please enter a temperature in Fahrenheit:")
	var fahr float64
	fmt.Scanf("%f", &fahr)
	fmt.Println(fahr)
	var celsius float64
	rightSide := 5 / 9
	fmt.Println(rightSide)
	celsius = (fahr - 32) * 5 / 9
	// celsius = 3.45
	fmt.Println(celsius)
	fmt.Println(reflect.TypeOf(celsius))
	fmt.Printf("Celsius %f", celsius)
}
