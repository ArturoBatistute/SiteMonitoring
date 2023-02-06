package log

import (
	"os"
	"strconv"
	"time"
)

func Register(site string, status bool) {

	const logFileName = "resources/output/log.txt"

	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		os.Create(logFileName)
	}

	var timestamp string = time.Now().Format("02/01/2006 15:04:05")

	file.WriteString(timestamp + " " + site + " online:" + strconv.FormatBool(status) + "\n")
	file.Close()
}
