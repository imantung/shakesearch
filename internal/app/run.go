package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

type (
	// Config ...
	Config struct {
		Port         string `default:"3001"`
		PreviewLimit int    `default:"500"`
		TextSource   string `default:"data/completeworks.txt"`
		MetaSource   string `default:"data/completeworks-meta.json"`
		Static       string `default:"./static"`
	}
)

// Run application
func Run() error {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	searcher, err := CreateSearcher(cfg)
	if err != nil {
		return err
	}

	e := echo.New()
	e.HideBanner = true
	e.Static("/", cfg.Static)
	e.GET("/search", func(ec echo.Context) error {
		q := ec.QueryParam("q")
		results := searcher.Search(q)
		return ec.JSON(http.StatusOK, results)
	})
	return e.Start(":" + cfg.Port)
}

// CreateSearcher ...
func CreateSearcher(cfg Config) (Searcher, error) {
	meta, err := readMeta(cfg.MetaSource)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(cfg.TextSource)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	book := NewBook(file, meta, cfg.PreviewLimit)
	return NewSuffixArraySearcher(book), nil
}

func readMeta(source string) (*Meta, error) {
	bytesMeta, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}
	var meta Meta
	if err := json.Unmarshal(bytesMeta, &meta); err != nil {
		return nil, err
	}
	return &meta, err
}
