package main

import "fmt"

func main() {
  fmt.Println(firstFactorial(6))
  fmt.Println(firstFactorial(4))
  fmt.Println(firstFactorial(8))
}

func firstFactorial(num int) int {
  factorial := 1
  for i := 1; i <= num; i++ {
    factorial = i * factorial
  }
  return factorial
}

