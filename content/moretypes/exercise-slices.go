// +build no-build OMIT

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var grade=make([][]uint8,dy)
	for y:=0; y<dy; y++ {
		grade[y] = make([]uint8, dx)
		for x:=0; x<dx; x++ {
			grade[y][x] = uint8((x+y)/2)
		}
	}
	return grade
}

func main() {
	pic.Show(Pic)
}
