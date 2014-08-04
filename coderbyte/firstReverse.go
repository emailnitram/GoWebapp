package main

import "fmt"

func main() {
  firstReverse("My name is Martin")
  firstReverse("Today is Sunday")
}

func firstReverse(s string) {
  fmt.Println(s.split())
}

