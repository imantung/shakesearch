package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Run application
func Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	searcher, err := NewSuffixArraySearcher("completeworks.txt")
	if err != nil {
		return err
	}

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
