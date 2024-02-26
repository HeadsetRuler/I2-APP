package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	reverse := flag.Bool("r", false, "Reverse lookup")
	flag.Parse()
	result, _ := dns(*reverse, flag.Arg(0))
	for _, r := range result {
		fmt.Println(r)
	}
}

func dns(reverse bool, address string) (result []string, err error) {
	if reverse {
		result, err = net.LookupAddr(address)
	} else {
		result, err = net.LookupHost(address)
	}
	return
}
