// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

type MyStruct struct {
	a string
	b string
}
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func (my MyStruct) M() {
	fmt.Println(my.a)
	fmt.Println(my.b)
}
func main() {
	var i I = T{"hello"}
	i.M()
	i=MyStruct{"hello", "world"}
	i.M()
}
