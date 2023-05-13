package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	htmlFileName := flag.String("html", "ex1.html", "html file name to be parsed")
	flag.Parse()

	htmlFile, err := os.Open(*htmlFileName)
	if err != nil {
		log.Fatalf("Couldn't open file %s %v ", *htmlFileName, err)
	}
	defer htmlFile.Close()

	doc, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatalf("Couldn't parse HTML %v ", err)
	}
	fmt.Println(doc)
}
