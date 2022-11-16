package main

import (
	"fmt"
)

func main() {
	// let use enter two numbers
	// sum all numbers between them including them
	a := 0
	b := 0

	fmt.Print("Enter a: ")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		return
	}

	fmt.Print("Enter b: ")
	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		return
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}
