package crypt

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"strconv"

	"github.com/ranon-rat/decChan/core"
)

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

func VerifySignature(signature, msgHash []byte, pubKey *rsa.PublicKey) bool {
	// If we don't get any error from the `VerifyPSS` method, that means our
	// signature is valid
	return rsa.VerifyPSS(pubKey, crypto.SHA256, msgHash, signature, nil) == nil
}
func GenHashPost(post core.Post) []byte {
	return sha256.New().
		Sum([]byte(post.User +
			post.Post +
			post.Title +
			strconv.Itoa(post.Date)))
}
func GenHashDelete(blockDeletion core.BlockDeletion) []byte {
	return sha256.New().Sum([]byte("DELETE" +
		blockDeletion.HashPost +
		strconv.Itoa(blockDeletion.DateDeletion) +
		strconv.Itoa(blockDeletion.DatePost)))

}
