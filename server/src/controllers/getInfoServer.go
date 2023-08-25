package controllers

import (
	"net/http"

	"github.com/ranon-rat/decChan/crypt"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	crypt.SendPubKey(PublicKey, w)
}
