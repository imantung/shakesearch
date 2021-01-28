package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type (
	// Config ...
	Config struct {
		Port         string `default:"3001"`
		PreviewLimit int    `default:"500"`
		Source       string `default:"completeworks.txt"`
	}
)

// Run application
func Run() error {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(cfg.Source)
	if err != nil {
		return err
	}

	searcher := NewSuffixArraySearcher(data, cfg.PreviewLimit)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/search", handleSearch(searcher))

	fmt.Printf("Listening on port %s...", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
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
