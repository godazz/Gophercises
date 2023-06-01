package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func findAnchors(node *html.Node, aChan chan *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		aChan <- node
		return
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		findAnchors(c, aChan)
	}

	if node.Parent == nil {
		close(aChan)
	}
}

func extractHref(aNode *html.Node) string {
	for _, attr := range aNode.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func extractText(aNode *html.Node) string {
	var text string
	for c := aNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
			continue
		}
		text += extractText(c)
	}

	return strings.TrimSpace(text)
}

func main() {
	htmlFileName := flag.String("html", "ex3.html", "html file name to be parsed")
	flag.Parse()

	htmlFile, err := os.Open(*htmlFileName)
	if err != nil {
		log.Fatalf("Couldn't open file %s %v ", *htmlFileName, err)
	}
	defer htmlFile.Close()

	root, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatalf("Couldn't parse HTML %v ", err)
	}

	aChan := make(chan *html.Node)
	go findAnchors(root, aChan)
	for a := range aChan {
		fmt.Println(Link{
			Href: extractHref(a),
			Text: extractText(a),
		})
	}
}
