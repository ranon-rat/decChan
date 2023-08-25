package controllers

import (
	"crypto/rsa"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
)

var (
	listConns  = make(map[core.ConnIP]bool)
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
	// later i will check on rsa for adding the public key and stuff

)

func SetupRSA() {
	PrivateKey, PublicKey = crypt.Init()
}
