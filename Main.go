package main

import "fmt"
import "net/http"
import "os"

func main(){

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
		monitoring()
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

func monitoring(){

	println("Monitoring..")

	siteUrlList := []string {"https://www.google.com", "https://www.linkedin.com"}
	
	for _,site := range siteUrlList{

		response,_ := http.Get(site)

		if response.StatusCode == 200 {
			fmt.Println(site, "is Up!")
		}else {
			fmt.Println(site, "is Down, returned status:", response.StatusCode)
		}
	}

	fmt.Println()
}