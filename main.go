package main

import (
	"fmt"

	"interface_demo/ishape"
	"interface_demo/shapes"
)

func main() {
	fmt.Println("This is a simple interface test with modules...")

	s := shapes.Square{Length: 2, Width: 2, Height: 2}
	printArea(s)
	printVolume(s)

	c := shapes.Circle{Radius: 55}
	printArea(c)
	printVolume(c)
}

func printArea(s ishape.Shape) {
	fmt.Printf("area: %T :: %v\n", s, s.Area())
}

func printVolume(s ishape.Shape) {
	fmt.Printf("volume: %T :: %v\n", s, s.Volume())
}
