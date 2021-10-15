package main

import (
	"fmt"

	"github.com/notd5a-alt/aurelias/port"
)

func Init() {
	fmt.Println("Aurelias Network Scanning tool")
	fmt.Println("Scanning Ports")

}

func PrettifyData(state bool, address string) string {
	return fmt.Sprintf("%v === %v\t\t", address, state)
}

func PrintScannedPorts(data []string) {
	var counter int

	for _, d := range data {
		fmt.Print(d)

		if counter%4 == 0 {
			fmt.Println()
		}

		counter++
	}
}

func main() {
	Init()
	tcp := port.TCP1024Scan() // TCP scan
	udp := port.UDP1024Scan() // UDP scan
	total_scan := append(tcp, udp...)
	var results []string

	for _, ts := range total_scan {
		results = append(results, PrettifyData(ts.State, ts.Address))
	}

	// finally print all scanned ports
	PrintScannedPorts(results)

}
