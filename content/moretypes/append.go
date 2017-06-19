// +build OMIT

package main

import "fmt"

func main() {
	var s []int
	sx:=make([]byte, 5)
	sx = sx[2:4]
	printSliceString(sx)
	t:=make([]byte , len(sx), (cap(sx)+1)*2)
	printSliceString(t)
	n:=t[3:7]
	printSliceString(n)
	n=n[:cap(n)]
	printSliceString(n)
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
	s = append(s, 2, 3, 4)
	s[9]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceString(s []byte) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
