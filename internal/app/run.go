package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	defaultPort         = "3001"
	defaultPreviewLimit = 500
	defaultSource       = "completeworks.txt"
)

// Run application
func Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	data, err := ioutil.ReadFile(defaultSource)
	if err != nil {
		return fmt.Errorf("Missing source: %w", err)
	}
	strings.Contains("", "")

	searcher := NewSuffixArraySearcher(data, defaultPreviewLimit)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/search", handleSearch(searcher))

	fmt.Printf("Listening on port %s...", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results := searcher.Search(query[0])
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}
