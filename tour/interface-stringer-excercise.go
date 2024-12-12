package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
type PrintIP struct {
	Name string
	IP   IPAddr
}

func (p PrintIP) String() string {
	return fmt.Sprintf("IP %v: %d.%d.%d.%d", p.Name, p.IP[0], p.IP[1], p.IP[2], p.IP[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		a := PrintIP{name, ip}
		fmt.Println(a)
	}
}
