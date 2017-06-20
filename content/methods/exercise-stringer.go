// +build OMIT

package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string  {
	return fmt.Sprintf("%[1]v.%[2]v.%[3]v.%[4]v",ip[0],ip[1],ip[2],ip[3])

}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	var a fmt.Stringer
	a=hosts["googleDNS"]
	fmt.Println(a)
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
