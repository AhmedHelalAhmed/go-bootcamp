package main

import (
	"fmt"
)

func main() {
	// first program
	fmt.Println("Hello, world!")

	// make it blue
	color := "green"
	color = "blue"
	fmt.Println(color)

	// vars to vars
	color = "green"
	color = "dark " + color
	fmt.Println(color)

	// assign with expressions
	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	// find the rectangle perimeter
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	// multi assign
	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	// multi assign #2
	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 19.5
	println("Air is good on " + planet)
	println("It's", isTrue)
	println("It's", temp, "degrees")

}
