package main

import (
	"log"

	"pulley.com/shakesearch/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
