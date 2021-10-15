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

func TCPLScan() []ScanResult {

	var results []ScanResult

	for i := 0; i <= 1024; i++ { // scans all ports from 0 -> 1024 i.e. :1 -> :2 -> :3 -> :4 and so on till :1024
		results = append(results, ScanPort("tcp", i))
	}

	return results

}

func UDPLScan() []ScanResult {

	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("udp", i))
	}

	return results

}

func UDPWScan() []ScanResult {

	var results []ScanResult

	for i := 1025; i <= 49152; i++ {
		results = append(results, ScanPort("udp", i))
	}

	return results

}

func TCPWScan() []ScanResult {

	var results []ScanResult

	for i := 1025; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", i))
	}

	return results
}

func UDPFullScan() []ScanResult {

	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("udp", i))
	}

	return results
}

func TCPFullScan() []ScanResult {

	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", i))
	}

	return results
}
