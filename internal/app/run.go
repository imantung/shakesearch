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

	book, err := CreateBook(cfg.Source, []string{
		"THE SONNETS",
		"ALL’S WELL THAT ENDS WELL",
		"THE TRAGEDY OF ANTONY AND CLEOPATRA",
		"AS YOU LIKE IT",
		"THE COMEDY OF ERRORS",
		"THE TRAGEDY OF CORIOLANUS",
		"CYMBELINE",
		"THE TRAGEDY OF HAMLET, PRINCE OF DENMARK",
		"THE FIRST PART OF KING HENRY THE FOURTH",
		"THE SECOND PART OF KING HENRY THE FOURTH",
		"THE LIFE OF KING HENRY THE FIFTH",
		"THE FIRST PART OF HENRY THE SIXTH",
		"THE SECOND PART OF KING HENRY THE SIXTH",
		"THE THIRD PART OF KING HENRY THE SIXTH",
		"KING HENRY THE EIGHTH",
		"KING JOHN",
		"THE TRAGEDY OF JULIUS CAESAR",
		"THE TRAGEDY OF KING LEAR",
		"LOVE’S LABOUR’S LOST",
		"THE TRAGEDY OF MACBETH",
		"MEASURE FOR MEASURE",
		"THE MERCHANT OF VENICE",
		"THE MERRY WIVES OF WINDSOR",
		"A MIDSUMMER NIGHT’S DREAM",
		"MUCH ADO ABOUT NOTHING",
		"THE TRAGEDY OF OTHELLO, MOOR OF VENICE",
		"PERICLES, PRINCE OF TYRE",
		"KING RICHARD THE SECOND",
		"KING RICHARD THE THIRD",
		"THE TRAGEDY OF ROMEO AND JULIET",
		"THE TAMING OF THE SHREW",
		"THE TEMPEST",
		"THE LIFE OF TIMON OF ATHENS",
		"THE TRAGEDY OF TITUS ANDRONICUS",
		"THE HISTORY OF TROILUS AND CRESSIDA",
		"TWELFTH NIGHT; OR, WHAT YOU WILL",
		"THE TWO GENTLEMEN OF VERONA",
		"THE TWO NOBLE KINSMEN",
		"THE WINTER’S TALE",
		"A LOVER’S COMPLAINT",
		"THE PASSIONATE PILGRIM",
		"THE PHOENIX AND THE TURTLE",
		"THE RAPE OF LUCRECE",
		"VENUS AND ADONIS",
	}, cfg.PreviewLimit)
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
