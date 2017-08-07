package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Printf("Need url argument at least.\n")
	} else {
		url := args[0]
		fmt.Printf("Running parse on: %s\n", url)
		parse(url)
	}
}

func parse(purl string) string {
	url := "https://mercury.postlight.com/parser?url=" + purl

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	MercuryApiKey := os.Getenv("MERCURY_API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "ERROR: GET REQUEST"
	}

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("x-api-key", MercuryApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "ERROR: CLIENT DO REQUEST"
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	sbody := string(body)

	fmt.Println(res)
	fmt.Println(sbody)

	return sbody
}
