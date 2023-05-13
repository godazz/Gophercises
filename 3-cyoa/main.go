package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
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

type StoryMux struct {
	story Story
}

func (m StoryMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storyArc := r.URL.Query().Get("arc")
	if storyArc == "" {
		storyArc = "intro" // start begining of story
	}

	arc, ok := m.story[storyArc]
	if !ok {
		http.Error(w, fmt.Sprintf("arc not found %s", storyArc), http.StatusNotFound)
		return
	}

	temp, err := template.ParseFiles("arc.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse template, %v", err), http.StatusInternalServerError)
	}

	if err := temp.Execute(w, arc); err != nil {
		http.Error(w, fmt.Sprintf("failed to execute template, %v", err), http.StatusInternalServerError)
	}
}

func NewstoryMux(story Story) http.Handler {
	return StoryMux{story}

}

func main() {
	jsonFileName := flag.String("json", "gopher.json", "path of a json file that represents the story")
	flag.Parse()

	file, err := os.Open(*jsonFileName)

	if err != nil {
		fmt.Printf("Failed to open %s: %v", *jsonFileName, err)
	}
	defer file.Close()

	var story Story
	if err := json.NewDecoder(file).Decode(&story); err != nil {
		fmt.Printf("Failed to decode %s: %v", *jsonFileName, err)
	}

	fmt.Printf("%+v\n", story)

	mux := NewstoryMux(story)
	http.ListenAndServe(":8080", mux)
}
