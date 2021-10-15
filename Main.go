package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/notd5a-alt/aurelias/port"
	"github.com/pterm/pterm"
)

const second = time.Second

func Init() int { // Prints inital ui etc for selecting modes etc.
	// ui below

	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("A", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("URELIAS", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"A Simple CLI Networking tool")

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nThis program was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()

	introSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Waiting for 5 seconds...")
	time.Sleep(second)
	for i := 5; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
		} else {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
		}
		time.Sleep(second)
	}
	introSpinner.Stop()

	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "0: Exit Program"},
		{Level: 0, Text: "Lower Ranged Scans, Ports: 0 -> 1024"},
		{Level: 1, Text: "1: TCP"},
		{Level: 1, Text: "2: UDP"},
		{Level: 1, Text: "3: TCP & UDP"},
		{Level: 0, Text: "Upper Ranged Scans, Ports: 1025 -> 49152"},
		{Level: 1, Text: "4: TCP"},
		{Level: 1, Text: "5: UDP"},
		{Level: 1, Text: "6: TCP & UDP"},
		{Level: 0, Text: "Full Range Scans, Ports: 0 -> 49152"},
		{Level: 1, Text: "7: TCP"},
		{Level: 1, Text: "8: UDP"},
		{Level: 1, Text: "9: TCP & UDP"},
	}).Render()
	pterm.DefaultSection.Println("Your Selection!")
	fmt.Print("= ")
	var input int
	fmt.Scanln(&input)

	return input

}

func PrettifyData(state bool, address string) string {
	return fmt.Sprintf("----------\t\t%v === %v\t\t", address, state)
}

func PrintScannedPorts(data *[]port.ScanResult) {
	d := pterm.TableData{{"Address", "State"}}
	for _, dd := range *data {
		if dd.State {
			d = append(d, []string{dd.Address, pterm.LightGreen("Pass")})
		} else {
			d = append(d, []string{pterm.LightRed(dd.Address), pterm.LightRed("Fail")})
		}
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
}

func main() {

	var scanResult []port.ScanResult
	var x bool = true

	for x {

		res := Init() // any initial print functions, ui, etc

		switch res {
		case 0:
			x = false
			os.Exit(0)
		case 1:
			scanResult = port.TCPLScan()
			x = false
		case 2:
			scanResult = port.UDPLScan()
			x = false
		case 3:
			x = false
			scanResult = append(port.TCPLScan(), port.UDPLScan()...)
		case 4:
			x = false
			scanResult = port.TCPWScan()
		case 5:
			x = false
			scanResult = port.UDPWScan()
		case 6:
			x = false
			scanResult = append(port.TCPWScan(), port.UDPWScan()...)
		case 7:
			x = false
			scanResult = port.TCPFullScan()
		case 8:
			x = false
			scanResult = port.UDPFullScan()
		case 9:
			x = false
			scanResult = append(port.TCPFullScan(), port.UDPFullScan()...)
		default:
			fmt.Println("Did not pick from the available choices, please pick again")
			fmt.Printf("\n\n\n\n")
		}
	}

	// finally print all scanned ports
	PrintScannedPorts(&scanResult)

}
