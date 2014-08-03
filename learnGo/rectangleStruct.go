package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	rect := Rectangle{4, 9}
	fmt.Println("Rectangle: Width: ", rect.width, " Height: ", rect.height)
	fmt.Println("Area", rect.Area())
	fmt.Println("Area", rect.Perimeter())
}
