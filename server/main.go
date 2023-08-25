package main

import (
	"os"

	"github.com/ranon-rat/decChan/server/src/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Setup(port)
}
