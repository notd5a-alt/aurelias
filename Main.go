package main

import (
	"fmt"
	"os"

	"github.com/notd5a-alt/aurelias/port"
)

func Init() int { // Prints inital ui etc for selecting modes etc.
	// ui below

	asciiArt :=
		`
		  _             _                  _           _            _              _          _                  _        
        / /\          /\_\               /\ \        /\ \         _\ \           /\ \       / /\               / /\      
       / /  \        / / /         _    /  \ \      /  \ \       /\__ \          \ \ \     / /  \             / /  \     
      / / /\ \       \ \ \__      /\_\ / /\ \ \    / /\ \ \     / /_ \_\         /\ \_\   / / /\ \           / / /\ \__  
     / / /\ \ \       \ \___\    / / // / /\ \_\  / / /\ \_\   / / /\/_/        / /\/_/  / / /\ \ \         / / /\ \___\ 
    / / /  \ \ \       \__  /   / / // / /_/ / / / /_/_ \/_/  / / /            / / /    / / /  \ \ \        \ \ \ \/___/ 
   / / /___/ /\ \      / / /   / / // / /__\/ / / /____/\    / / /            / / /    / / /___/ /\ \        \ \ \       
  / / /_____/ /\ \    / / /   / / // / /_____/ / /\____\/   / / / ____       / / /    / / /_____/ /\ \   _    \ \ \      
 / /_________/\ \ \  / / /___/ / // / /\ \ \  / / /______  / /_/_/ ___/\ ___/ / /__  / /_________/\ \ \ /_/\__/ / /      
/ / /_       __\ \_\/ / /____\/ // / /  \ \ \/ / /_______\/_______/\__\//\__\/_/___\/ / /_       __\ \_\\ \/___/ /       
\_\___\     /____/_/\/_________/ \/_/    \_\/\/__________/\_______\/    \/_________/\_\___\     /____/_/ \_____\/              
		`

	asciiArt2 :=
		`
		  _            _        _          _          _            _             _        
        /\ \         /\ \     /\ \       /\ \       /\ \         /\ \     _    / /\      
       /  \ \       /  \ \    \_\ \      \ \ \     /  \ \       /  \ \   /\_\ / /  \     
      / /\ \ \     / /\ \ \   /\__ \     /\ \_\   / /\ \ \     / /\ \ \_/ / // / /\ \__  
     / / /\ \ \   / / /\ \_\ / /_ \ \   / /\/_/  / / /\ \ \   / / /\ \___/ // / /\ \___\ 
    / / /  \ \_\ / / /_/ / // / /\ \ \ / / /    / / /  \ \_\ / / /  \/____/ \ \ \ \/___/ 
   / / /   / / // / /__\/ // / /  \/_// / /    / / /   / / // / /    / / /   \ \ \       
  / / /   / / // / /_____// / /      / / /    / / /   / / // / /    / / /_    \ \ \      
 / / /___/ / // / /      / / /   ___/ / /__  / / /___/ / // / /    / / //_/\__/ / /      
/ / /____\/ // / /      /_/ /   /\__\/_/___\/ / /____\/ // / /    / / / \ \/___/ /       
\/_________/ \/_/       \_\/    \/_________/\/_________/ \/_/     \/_/   \_____\/       
	`

	fmt.Println("Thanks for using")
	fmt.Println()

	fmt.Println(asciiArt)
	fmt.Printf("\n\n")
	fmt.Println("A Simple networking tool written in GoLang")
	fmt.Println()

	fmt.Println(asciiArt2)
	fmt.Printf("\n\n")
	fmt.Println("Scan Mode: ")
	fmt.Println("0. Exit Program")
	fmt.Println("1. Lower Scan UDP")
	fmt.Println("2. Lower Scan TCP")
	fmt.Println("3. UDP & TCP")
	fmt.Println("4. Wide Scan UDP")
	fmt.Println("5. Wide Scan TCP")
	fmt.Println("6. Wide Scan UDP & TCP")
	fmt.Println("7. Full Scan UDP")
	fmt.Println("8. Full Scan TCP")
	fmt.Println("9. Full Scan UDP & TCP")
	fmt.Print("Your Selection === ")
	var input int
	fmt.Scanln(&input)

	return input

}

func PrettifyData(state bool, address string) string {
	return fmt.Sprintf("%v === %v\t\t", address, state)
}

func PrintScannedPorts(data *[]port.ScanResult) {
	var counter int

	for _, d := range *data {
		fmt.Print(PrettifyData(d.State, d.Address))

		if counter%4 == 0 { // prints 4 pieces of data tabbed
			fmt.Println()
		}

		counter++
	}
}

func main() { // TODO: Clean up main function
	res := Init() // any initial print functions, ui, etc

	var scanResult []port.ScanResult
	var x bool = true

	for x {

		switch res {
		case 0:
			x = false
			os.Exit(0)
		case 1:
			scanResult = port.TCP1024Scan()
			x = false
		case 2:
			scanResult = port.UDP1024Scan()
			x = false
		case 3:
			x = false
			scanResult = append(port.UDP1024Scan(), port.TCP1024Scan()...)
		case 4:
			scanResult = port.UDP49152Scan()
		case 5:
			scanResult = port.TCP49152Scan()
		case 6:
			scanResult = append(port.UDP49152Scan(), port.TCP49152Scan()...)
		case 7:
			scanResult = append(port.UDP1024Scan(), port.UDP49152Scan()...)
		case 8:
			scanResult = append(port.TCP1024Scan(), port.TCP49152Scan()...)
		case 9:
			scanResult = append(port.UDP1024Scan(), port.UDP49152Scan()...)
			scanResult = append(scanResult, port.TCP1024Scan()...)
			scanResult = append(scanResult, port.TCP49152Scan()...)
		default:
			fmt.Println("Did not pick from the available choices, please pick again")
			fmt.Printf("\n\n\n\n")
		}

	}

	// finally print all scanned ports
	PrintScannedPorts(&scanResult)

}
