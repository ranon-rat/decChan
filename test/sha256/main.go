package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenHash(msg string) []byte {
	msgHash := sha256.New()
	msgHash.Write([]byte(msg))

	return msgHash.Sum(nil)
}
func main() {
	a := "fuck fuck fuck"
	b := "culo carajo"
	fmt.Println("a", hex.EncodeToString(GenHash(a)))
	fmt.Println("b", hex.EncodeToString(GenHash(b)))
}
