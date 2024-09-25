package main

import (
	"github.com/olga-sinepalnikova/creativemobile-testtask/internal/config"
	"log"
)

func main() {
	// todo: use logrus
	_, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
}
