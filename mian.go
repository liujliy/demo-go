package main

import (
	"fmt"
	"image/color"
)

func main() {
	fmt.Println("Hello Go!")

	// struct
	test := ColorPoint{
		Point: Point{3, 4},
		Color: color.RGBA{G: 225, B: 225, A: 1},
	}
	fmt.Println(test)
}

type Point struct {
	x, y float64
}

type ColorPoint struct {
	Point
	Color color.RGBA
}
