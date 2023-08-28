package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	_, err := http.Post("http://localhost:9090", "", nil)
	fmt.Println(err.Error(), len(err.Error()))
	fmt.Println(strings.Contains("shitty bido", "bido"))
}
