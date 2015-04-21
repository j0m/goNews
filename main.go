package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"./feedParser"
)

func main() {
	response, err := http.Get("http://blog.fefe.de/rss.xml")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		//fmt.Printf("%s\n", string(contents))

		rss := feedParser.NewRssFeed()
		rss.Parse(contents)
    }
}
