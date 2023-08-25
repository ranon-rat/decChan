package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"
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

func SignMSG(priKey *rsa.PrivateKey, msgHashSum []byte) (signature []byte) {

	// i dont think that there is no need to save the hash, but i will save it in the server
	signature, _ = rsa.SignPSS(rand.Reader, priKey, crypto.SHA256, msgHashSum, nil)
	return

}
