package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/ranon-rat/decChan/core"
)

func Init() (privKey *rsa.PrivateKey, pubKey *rsa.PublicKey) {
	privKey, pubKey = ParseKeys("private.pem")
	if privKey == nil {
		privKey, pubKey = SaveKey()
	}
	return
}

// this will return you the public and private key, this is oriented for the server
func ParseKeys(namefile string) (prvKey *rsa.PrivateKey, pubKey *rsa.PublicKey) {
	pemFile, err := os.ReadFile(namefile)
	if err != nil {
		log.Print(core.ErrMsg, "it appears that the", namefile, "doesnt exists")
		return
	}
	block, _ := pem.Decode(pemFile)

	prvKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// i dont need to end the process here, because if there is any error i would end up the process inmediatly
		// so this is importante
		log.Println(core.ErrMsg, " it appears that something is wrong with the file", namefile, "generate a new one")

	}
	pubKey = &prvKey.PublicKey

	return
}

/*
this is for the server

	this is conf stuff and if it doesnt works the network is not  secure, so the service will stop if something is wrong
*/
func SaveKey() (privKey *rsa.PrivateKey, pubKey *rsa.PublicKey) {
	// there is no way that this could generate a problem
	privKey, _ = rsa.GenerateKey(rand.Reader, 2048)

	pubKey = &privKey.PublicKey

	// save PEM file
	pemfile, err := os.Create("private.pem")
	if err != nil {
		log.Panic(core.ErrMsg, err)

	}
	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey)}
	// since im saving it to a file i dont think that there is some need to check if there is any error
	pem.Encode(pemfile, pemkey)
	pemfile.Close()
	return
}
