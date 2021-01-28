package app

import (
	"io/ioutil"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

type (
	// Config ...
	Config struct {
		Port         string `default:"3001"`
		PreviewLimit int    `default:"500"`
		Source       string `default:"completeworks.txt"`
		Static       string `default:"./static"`
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
