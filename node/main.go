package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/ranon-rat/decChan/node/src/router"
)

var port = 8000 + rand.Intn(100)

func main() {
	router.Setup(strconv.Itoa(port))
	fmt.Println("this isnt ready i need to finish the server, wait a few days for ending this")

}
