package monitoring

import (
	"fmt"
	"net/http"
	"time"
	"projects/UrlMonitoring/src/reader"
	"projects/UrlMonitoring/src/log"
)

func Run(timesToExecute int) {

	const delayUntilNextMonitoring = 5

	println("Monitoring..")

	for i := 0; i < timesToExecute; i++ {

		runRequest()
		time.Sleep(delayUntilNextMonitoring * time.Second)
	}

	fmt.Println()
}

func runRequest() {

	siteUrlList := reader.ReadSitesFromFile()

	for _, site := range siteUrlList {

		response, err := http.Get(site)

		if err != nil {
			fmt.Println(err)
		}

		if response.StatusCode == 200 {
			log.Register(site, true)
			fmt.Println(site, "is Up!")
		} else {
			log.Register(site, false)
			fmt.Println(site, "is Down, returned status:", response.StatusCode)
		}
	}
}
