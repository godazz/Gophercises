package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlStruct struct {
	path string `yaml:"path, inline"`
	url  string `yaml:"url, inline"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if longURL, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, longURL, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type yamlData struct {
	Path string
	Url  string
}

func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(data)
	fmt.Println(parsedYAML)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYAML)
	fmt.Println(pathMap)

	return MapHandler(pathMap, fallback), nil
}

func parseYAML(data []byte) ([]yamlData, error) {
	var parsedYAML []yamlData
	err := yaml.Unmarshal([]byte(data), &parsedYAML)
	if err != nil {
		return nil, err
	}
	return parsedYAML, nil
}

func buildMap(parsedYAML []yamlData) map[string]string {
	mp := make(map[string]string)

	for _, v := range parsedYAML {
		mp[v.Path] = v.Url
	}
	return mp
}
