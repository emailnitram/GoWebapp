package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/",fs)
}