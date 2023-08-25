package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	privKey, pubKey := ParseKeys("private.pem")
	if privKey == nil {
		privKey, pubKey = SaveKey()
	}
	msg := "hello world bitches"
	sign, hash := SignMSG(privKey, msg)

	fmt.Printf("sign %d,hash %d msg %s\n", len(fmt.Sprintf("%x", sign)), len(fmt.Sprintf("%x", hash)), msg)

	fmt.Println(VerifySignature(sign, hash, pubKey))
	fmt.Println(VerifySignature(sign, []byte{}, pubKey))
}

// this will return you the public and private key, this is oriented for the server
func ParseKeys(namefile string) (prvKey *rsa.PrivateKey, pubKey *rsa.PublicKey) {
	pemFile, err := os.ReadFile(namefile)
	if err != nil {
		log.Panic("[-] it appears that the", namefile, "doesnt exists")
		return
	}
	block, _ := pem.Decode(pemFile)

	prvKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// i dont need to end the process here, because if there is any error i would end up the process inmediatly
		// so this is importante
		log.Println("[-] it appears that something is wrong with the file", namefile, "generate a new one")

	}
	pubKey = &prvKey.PublicKey

	return
}
func SendPubKey(pubKey *rsa.PublicKey, w http.ResponseWriter) {
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pubKey),
	}
	pem.Encode(w, publicKeyPEM)
}
func ParsePubKey(pubKeyB []byte) (pubKey *rsa.PublicKey) {

	block, _ := pem.Decode(pubKeyB)

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		// i dont need to end the process here, because if there is any error i would end up the process inmediatly
		// so this is importante
		log.Println("[-] it appears that something is wrong with the public key, report this to the server admins https://discord.gg/yRjNapa4ud")

	}
	return pubKey
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
		log.Panic("[-]", err)

	}
	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey)}
	// since im saving it to a file i dont think that there is some need to check if there is any error
	pem.Encode(pemfile, pemkey)
	pemfile.Close()
	return
}

/*
	this is for the server

it will return a signature and the hashmap, i will be responsable for encoding it to hex or base64
also, the size of the signature is 512 if we are using hex for coding it, and the hash is of 64
so i need to keep that in mind while working in the database.
*/
func SignMSG(priKey *rsa.PrivateKey, msg string) (signature, msgHashSum []byte) {

	// i dont think that there is no need to save the hash, but i will save it in the server
	msgHashSum = sha256.New().Sum(nil)
	signature, _ = rsa.SignPSS(rand.Reader, priKey, crypto.SHA256, msgHashSum, nil)
	return

}

/*
	this is for the client

the client needs to have the public key
I need to send it and parse it
*/
func VerifySignature(signature, msgHash []byte, pubKey *rsa.PublicKey) bool {
	// If we don't get any error from the `VerifyPSS` method, that means our
	// signature is valid
	return rsa.VerifyPSS(pubKey, crypto.SHA256, msgHash, signature, nil) == nil
}
