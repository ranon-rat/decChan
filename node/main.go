package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/node/src/router"
)

var port = 8921

func main() {
	rand.Seed(time.Now().Unix())
	if core.Dev {
		port = 8000 + rand.Intn(100)
	}
	fmt.Println(`
	                                    
	____             _____ _           
	|    \ ___ ___   |     | |_ ___ ___ 
	|  |  | -_|  _|  |   --|   | .'|   |
	|____/|___|___|  |_____|_|_|__,|_|_|
										
	
	A project made with love by @bruh-boys
	
	Welcome to thisw 
	`)
	fmt.Println("\n\nI recomend you to go check http://localhost:" + strconv.Itoa(port))
	router.Setup(strconv.Itoa(port))

}
