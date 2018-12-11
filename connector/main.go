package main

import (
	"flag"
	"fmt"
)

func main() {
	var remote = flag.String("remote", "https://10.61.18.217:8081", "remote address")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: %s remote_URI username password ")
	}
}