// Download Research Papers (because information should be free for all)
// A wrapper around sci-hub.tw (credit: Alexandra Elbakyan)
// Author - Kumar Ashwin
// Version - v2.0

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fatih/color"
)

func init() {
	flag.Usage = func() {
		h := `
desearch v2.0
CLI Wrapper for sci-hub.tw, download research papers for free!
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

Usage:
	desearch [url]
	
Author:
	[Kumar Ashwin](krashwin00@gmail.com)
`
		fmt.Fprintf(os.Stderr, h)
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Print("Enter valid Arguments. Type `desearch -h` for help!\n")
		os.Exit(1)
	} else {

		flag.Parse()
		base := "https://sci-hub.tw/"
		research := os.Args[1]
		completeURL := base + research
		response, err := http.Get(completeURL)
		checkErr(err)
		defer response.Body.Close()

		dataInBytes, err := ioutil.ReadAll(response.Body)
		checkErr(err)
		pageContent := string(dataInBytes)

		linkIndex := strings.Index(pageContent, "?download=true")
		lastIndex := linkIndex
		startIndex := 0
		for true {
			if strings.Compare(string(pageContent[linkIndex]), "'") == 0 {
				startIndex = linkIndex
				break
			}
			linkIndex = linkIndex - 1
		}
		downloadLink := pageContent[startIndex+1 : lastIndex]

		if !strings.HasPrefix(downloadLink, "http") {
			downloadLink = "https:" + downloadLink
		}

		fmt.Print("Downloading Research Paper: ")

		// Building FileName from Download Link
		fileURL := url.Parse(downloadLink)
		path := fileURL.Path
		segment := strings.Split(path, "/")
		filename := segment[len(segment)-1]
		color.Cyan(filename)

		// Creating a file to put in data
		output, err := os.Create(filename)
		checkErr(err)
		defer output.Close()

		response, err = http.Get(downloadLink)
		checkErr(err)
		defer response.Body.Close()

		// Checking for response code to be OK
		if response.StatusCode != http.StatusOK {
			fmt.Printf("Cannot Download File: Status not OK!")
		}

		// Copy the data to the output file created.
		_, err = io.Copy(output, response.Body)
		checkErr(err)
		fmt.Print("Done\n")
	}
}

func checkErr(err error) {
	if err != nil {
		color.Red("Cannot download specified paper...please check manually @ sci-hub.tw\nSorry for the inconvenience!")
		os.Exit(1)
	}
}
