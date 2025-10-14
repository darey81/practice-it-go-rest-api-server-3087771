package main

import (
	"fmt"
)
func Volume(length, width, height int) int {
	return length * width * height
}

func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height
	return
}

func main() {
	area, volume := dimensions(5,5,5)
	fmt.Println("Area: ", area, "Volume: ", volume)
	v := Volume(10,11,12)
	fmt.Println("Volume: ", v)
}