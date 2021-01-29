package app

import (
	"errors"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

type (
	// Config ...
	Config struct {
		Port         string `default:"3001"`
		PreviewLimit int    `default:"500"`
		Source       string `default:"data/completeworks.txt"`
		Static       string `default:"./static"`
	}
)

// Run application
func Run() error {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	book, err := CreateBook(cfg.Source, []string{}, cfg.PreviewLimit)
	if err != nil {
		return errors.New("Book: " + err.Error())
	}

	searcher := NewSuffixArraySearcher(book)

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
