package main

//pointers v. value methods

import (
	"fmt"
)
type Dimension struct {
	length 	int
	width 	int
	height	int
}
func (d *Dimension) Area () int { // added * to test pointer
	d.height = 8
	return d.length * d.width
}
func (d Dimension) Volume() int {
	d.height = 6
	return d.length * d.width * d.height
}

type MyInt int

func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height
	return
}

func main() {
	d := Dimension{20,5,6}
	fmt.Println(d.Area())	
	fmt.Println(d)
	fmt.Println(d.Volume())
	fmt.Println(d)
}
