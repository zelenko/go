package main

import (
	"fmt"
	"strconv"
)

// IPAddr is the IPv4 address
type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// String implements stringer interface
func (p IPAddr) String() string {
	s := strconv.Itoa(int(p[0])) + "." +
		strconv.Itoa(int(p[1])) + "." +
		strconv.Itoa(int(p[2])) + "." +
		strconv.Itoa(int(p[3]))
	return `"` + s + `"`
}
