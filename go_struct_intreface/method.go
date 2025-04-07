package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Method with a valud receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with a pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	r := Rectangle{Width: 10, Height: 5}

	fmt.Println("Area:", r.Area()) // Output: Area: 50

	r.Scale(2)
	fmt.Println("After scaling - Width:", r.Width, "Height:", r.Height)
	fmt.Println("New Area:", r.Area())
}
