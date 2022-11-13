package main

import (
	"fmt"
	"os"
	"path"
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

	// multi short func
	_, b := multi()
	fmt.Println(b)

	// swapper
	color1, color2 := "red", "blue"
	color1, color2 = "orange", "green"
	fmt.Println(color1, color2)

	// swapper #2
	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)

	// discard the file
	dir, _ := path.Split("secret/file.txt")
	println(dir)

	// convert and fix #1
	aa, bb := 10, 5.5
	fmt.Println(float64(aa) + bb)

	// convert and fix #2
	aaa, bbb := 10, 5.5
	aaa = int(bbb)
	fmt.Println(float64(aaa) + bbb)

	// convert and fix #3
	fmt.Println(float32(5.5))

	// convert and fix #4
	age := 2
	fmt.Println(7.5 + float32(age))

	// convert and fix #5
	min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min))

	// count arguments
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)

	// print the path
	fmt.Println(os.Args[0])

	// print your name
	if len(os.Args)-1 >= 1 {
		fmt.Println(os.Args[1])
		fmt.Println("Hello", os.Args[1])
		fmt.Println("How are you?")
	}

	// greet more people

	peopleCount := len(os.Args) - 1
	people := os.Args
	if peopleCount > 0 {
		fmt.Println("There are", peopleCount, "people")
	}
	if peopleCount >= 1 {
		fmt.Println("Hello great", people[1], "!")
	}
	if peopleCount >= 2 {
		fmt.Println("Hello great", people[2], "!")
	}
	if peopleCount >= 3 {
		fmt.Println("Hello great", people[3], "!")
	}
	fmt.Println("Nice to meet you all.")

}

func multi() (int, int) {
	return 5, 4
}
