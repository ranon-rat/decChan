package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ranon-rat/decChan/node/src/router"
)

var port = 8000 + rand.Intn(100)

func main() {
	rand.Seed(time.Now().Unix())
	port = 8000 + rand.Intn(100)
	fmt.Println("this isnt ready i need to finish the server, wait a few days for ending this")
	fmt.Println("\n\nI recomend you to check http://localhost:" + strconv.Itoa(port))
	router.Setup(strconv.Itoa(port))

}
