package link

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Text string
	Href string
}

func Parse(siteURL string) []Link {
	htmlRoot := getPageHTML(siteURL)

	aChan := make(chan *html.Node)
	go findAnchors(htmlRoot, aChan)

	var links []Link
	for a := range aChan {
		l := Link{
			Text: extractText(a),
			Href: extractHref(a),
		}
		links = append(links, l)
	}
	return links
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

func getPageHTML(siteURL string) *html.Node {
	resp, err := http.Get(siteURL)
	if err != nil {
		log.Fatalf("Failed to Get http Response %v", err)
	}
	defer resp.Body.Close()
	htmlRoot, err := html.Parse(resp.Body)

	if err != nil {
		log.Fatalf("Failed to Parse html document %v", err)
	}

	return htmlRoot
}
