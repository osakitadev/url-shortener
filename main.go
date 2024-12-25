package main

import (
	"flag"
	"log"
)

var (
	URL string
)

func init() {
	flag.StringVar(&URL, "url", "", "URL to shorten")
	flag.Parse()
}

func main() {
	if URL == "" {
		panic("URL is required")
	}
	hashedURL := HashURL(URL)

	log.Println("Shortened URL:", "http://localhost:8080/"+hashedURL)

	SaveURLToFile(hashedURL, URL)
	ServeServer()
}
