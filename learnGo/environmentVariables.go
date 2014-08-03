package main

import (
	"fmt"
	"os"
	s "strings"
)

func main() {
	var keyValEnv string
	fmt.Println("Please enter an enviornment variable seperated by =")
	fmt.Scanf("%s", &keyValEnv)
	keyvalSlice := s.Split(keyValEnv, "=")
	os.Setenv(keyvalSlice[0], keyvalSlice[1])

	environmentVariables := os.Environ()
	for _, value := range environmentVariables {
		fmt.Println(value)
	}
}
