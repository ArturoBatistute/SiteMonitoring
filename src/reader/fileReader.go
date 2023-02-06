package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadSitesFromFile() []string {

	const sitesFileName = "resources/sites.txt"
	var siteSlice []string

	file, err := os.Open(sitesFileName)

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
