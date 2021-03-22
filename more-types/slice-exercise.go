package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	fmt.Printf("len=%d cap=%d %v\n", len(img), cap(img), img)

	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8(x ^ y)
		}
		img[y] = s
	}
	return img
}

func run_slice_exercise() {
	pic.Show(Pic)
}
