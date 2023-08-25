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
	fmt.Println("not ready, fuck you ")

}
