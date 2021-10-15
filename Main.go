package main

import (
	"fmt"

	"github.com/notd5a-alt/aurelias/port"
)

func init() {
	fmt.Println("Welcome to the ")
	fmt.Println("Scanning Ports")

}

func main() {
	results := port.TCPScan()

}
