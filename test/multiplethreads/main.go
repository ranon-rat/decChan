package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go doSomethingHere()
	http.ListenAndServe(":8080", nil)
}
func doSomethingHere() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Second * 2)
			fmt.Println(i, "hello")
		}(i)
	}
}
