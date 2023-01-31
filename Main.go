package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const delayUntilNextMonitoring = 5

func main(){

	readSitesFromFile()

	for {
		showMenu()
		userChoice()
	}
}

func showMenu(){

	fmt.Println("Choose an option:")
	fmt.Println("1- Monitoring")
	fmt.Println("2- Logs")
	fmt.Println("0- Exit")
	fmt.Println()
}

func userChoice(){

	var command int
	fmt.Scan(&command)

	fmt.Println()

	switch command{
	case 1:
		monitoring(3)
	case 2:
		println("Logging..")
	case 0:
		println("Exiting..")
		os.Exit(0)
	default:
		println("Please select an existing command.")
		os.Exit(-1)
	}
}

func monitoring(timesToExecute int){

	println("Monitoring..")

	for i := 0; i < timesToExecute; i++ {

		monitoringRun()
		time.Sleep(delayUntilNextMonitoring * time.Second)
	}

	fmt.Println()
}

func monitoringRun(){

	siteUrlList := readSitesFromFile()
	
	for _,site := range siteUrlList{

		response, err := http.Get(site)

		if err != nil {
			fmt.Println(err)
		}

		if response.StatusCode == 200 {
			fmt.Println(site, "is Up!")
		}else {
			fmt.Println(site, "is Down, returned status:", response.StatusCode)
		}
	}
}

func readSitesFromFile() []string {

	var siteSlice []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		siteSlice = append(siteSlice, strings.TrimSpace(line))

		if err == io.EOF {
			break
		}
	}
	file.Close()
	return siteSlice
}