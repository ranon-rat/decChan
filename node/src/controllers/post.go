package controllers

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
	"github.com/ranon-rat/decChan/node/src/db"
)

// deberia de trabajar en el formato para poder hacer esto mas sencillo
// luego lo agregare
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1/get-post?date=blah&board=blah
	// board can be
	date, err := strconv.Atoi(r.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, "non valid date", http.StatusBadRequest)
		return
	}
	board := r.URL.Query().Get("board")
	if board == "" {
		http.Error(w, "empty field", http.StatusBadRequest)
		return
	}
	if !core.Boards[board] && !db.CheckExistencePosts(board) {
		http.Error(w, "non existence of the board or the post", http.StatusBadRequest)
		return
	}
	blocks := db.GetPosts(board, date)
	core.PrintInfo(blocks)
	json.NewEncoder(w).Encode(blocks)
}
func NewPost(w http.ResponseWriter, r *http.Request) {

	var blockPost core.BlockPost
	if json.NewDecoder(r.Body).Decode(&blockPost) != nil {
		http.Error(w, "formato no valido", http.StatusBadRequest)
		return
	}
	post := blockPost.Post
	hash := crypt.GenHashPost(post)
	signature, err := hex.DecodeString(blockPost.Signature)

	if err != nil {
		core.PrintErr(err)
		return
	}
	core.PrintInfo(blockPost)
	// verify everything
	if !crypt.VerifySignature(signature, hash, pubKey) {
		core.PrintInfo("someone sent something weird")
		http.Error(w, "non valid block", 404)
		return
	}

	if !core.Boards[post.Board] && !db.CheckExistencePosts(post.Board) {
		core.PrintInfo("someone sent a non valid board delete it or be gay")
		return
	}
	if db.CheckExistencePosts(post.Board) && db.ItGotToLimit(post.Board) {
		core.PrintErr("this got to the limit")
		http.Error(w, "i cant accept more post from this thread, avoid it", http.StatusBadRequest)

		return
	}
	// tecnicamente no puedo repetir algo aqui asi que no deberia de haber problema, solamente
	// que alguien puede replicar esto
	// aun que de todas maneras lo detendria la base de datos

	hashS := hex.EncodeToString(hash)
	if db.CheckExistenceDeletion(hashS) || db.CheckExistencePosts(hashS) {
		core.PrintErr("this block is repeated")

		http.Error(w, "this is just trash, ignore it", http.StatusBadRequest)

		return
	}

	db.AddPost(blockPost)
	blocksChan <- BlockSender{
		Blocks: core.Blocks{
			BlocksPosts: []core.BlockPost{blockPost},
		}}
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
	if db.CheckExistenceDeletion(blockDel.HashPost) {
		return
	}
	blocksChan <- BlockSender{
		Blocks: core.Blocks{
			BlocksDeletion: []core.BlockDeletion{blockDel},
		},
	}
	db.DeletePost(blockDel)
}
