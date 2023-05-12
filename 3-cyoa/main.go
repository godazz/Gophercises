package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Story map[string]struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {
	jsonFileName := flag.String("json", "gopher.json", "path of a json file that represents the story")
	flag.Parse()

	file, err := os.Open(*jsonFileName)

	if err != nil {
		fmt.Printf("Failed to open %s: %v", *jsonFileName, err)
		return
	}
	defer file.Close()

	var story Story
	if err := json.NewDecoder(file).Decode(&story); err != nil {
		fmt.Printf("Failed to decode %s: %v", *jsonFileName, err)
		return
	}

	fmt.Printf("%+v\n", story)

}
