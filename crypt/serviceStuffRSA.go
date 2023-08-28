package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"

	"github.com/ranon-rat/decChan/core"
)

// this is for the controllers
func SendPubKey(pubKey *rsa.PublicKey, w http.ResponseWriter) {
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pubKey),
	}
	pem.Encode(w, publicKeyPEM)
}

/*
	this is for the server

it will return a signature and the hashmap, i will be responsable for encoding it to hex or base64
also, the size of the signature is 512 if we are using hex for coding it, and the hash is of 64
so i need to keep that in mind while working in the database.
*/

func SignMSG(priKey *rsa.PrivateKey, hashSum []byte) (signature []byte) {
	signature, err := rsa.SignPSS(rand.Reader, priKey, crypto.SHA256, hashSum, nil)
	if err != nil {
		core.PrintErr(err)
	}
	return

}
