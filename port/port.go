package port

import (
	"net"
	"strconv"
	"time"
)

// Struct to hold scan data
type ScanResult struct {
	Port    int
	State   bool
	Address string
}

// Scanning functions

func ScanPort(protocol string, port int) ScanResult {
	var result ScanResult = ScanResult{Port: port}
	var address string = ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 30*time.Second)

	// error checking
	if err != nil {
		result.State = false
		result.Address = address
		return result
	}

	defer conn.Close()
	result.State = true
	result.Address = address

	return result
}

func TCP1024Scan() []ScanResult {

	var results []ScanResult

	for i := 0; i <= 1024; i++ { // scans all ports from 0 -> 1024 i.e. :1 -> :2 -> :3 -> :4 and so on till :1024
		results = append(results, ScanPort("tcp", i))
	}

	return results

}

func UDP1024Scan() []ScanResult {

	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("udp", i))
	}

	return results

}
