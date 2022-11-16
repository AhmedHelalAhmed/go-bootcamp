package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	x := Vertex{1.5, 1.5}
	fmt.Println(Vertex{1.5, 1.5})
	fmt.Println(x)
	// struct itself
	x.Scale(10)
	fmt.Println(x)
	ScaleFunc(&x, 10)
	fmt.Println(x)

	y := &x

	// address of the struct
	(&x).Scale(10)
	y.Scale(10)
	fmt.Println(*y)
	ScaleFunc(y, 10)
}
