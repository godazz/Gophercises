package main

import (
	"flag"
	"fmt"

	"github.com/godazz/Gophercises/sitemap/link"
)

func main() {
	siteURL := flag.String("url", "http://example.com/", "The url for the website to build the sitemap for.")
	flag.Parse()

	links := link.Parse(*siteURL)

	for _, link := range links {
		fmt.Println(link.Text, link.Href)
	}

}
