package controllers

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
	"github.com/ranon-rat/decChan/node/src/db"
)

func NewPost(w http.ResponseWriter, r *http.Request) {

	var blockPost core.BlockPost
	if json.NewDecoder(r.Body).Decode(&blockPost) != nil {
		http.Error(w, "formato no valido", http.StatusBadRequest)
		return
	}
	// simple crypto cool stuff
	hash := crypt.GenHashPost(blockPost.Post)
	signature, err := hex.DecodeString(blockPost.Signature)
	if err != nil {
		return
	}
	// verify everything
	if !crypt.VerifySignature(signature, hash, pubKey) {
		core.PrintInfo("someone sent something weird")
		return
	}
	blocksChan <- BlockSender{
		Blocks: core.Blocks{
			BlocksPosts: []core.BlockPost{blockPost},
		}}
	hashS := hex.EncodeToString(hash)
	if db.CheckExistenceDeletion(hashS) || db.CheckExistencePosts(hashS) {
		return
	}
	db.AddPost(blockPost)

}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	var blockDel core.BlockDeletion
	if json.NewDecoder(r.Body).Decode(&blockDel) != nil {
		http.Error(w, "formato no valido", http.StatusBadRequest)
		return
	}
	hash := crypt.GenHashDelete(blockDel)
	signature, err := hex.DecodeString(blockDel.Signature)
	if err != nil {
		return
	}
	if !crypt.VerifySignature(signature, hash, pubKey) {
		core.PrintInfo("someone sent something weird")
	}
	blocksChan <- BlockSender{
		Blocks: core.Blocks{
			BlocksDeletion: []core.BlockDeletion{blockDel},
		},
	}
	db.DeletePost(blockDel)
}
