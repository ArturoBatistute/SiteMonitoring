package main

import (
	"projects/UrlMonitoring/src/monitoring"
	"fmt"
	"os"
)

func main() {

	for {
		showMenu()
		userChoice()
	}
}

func showMenu() {

	fmt.Println("Choose an option:")
	fmt.Println("1- Monitoring")
	fmt.Println("0- Exit")
	fmt.Println()
}

func userChoice() {

	var command int
	fmt.Scan(&command)

	fmt.Println()

	switch command {
	case 1:
		monitoring.Run(3)
	case 0:
		println("Exiting..")
		os.Exit(0)
	default:
		println("Please select an existing command.")
		os.Exit(-1)
	}
}