// +build no-build OMIT

package main

import "github.com/Go-zh/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (my_r MyReader) Read(b []byte) (int, error)  {
	my_r.

}

func main() {
	reader.Validate(MyReader{})
}
