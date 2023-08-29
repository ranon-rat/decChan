package controllers

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
)

func Board(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetBoard(w, r)
	case "POST":
		PostBoard(w, r)
	case "DELETE":
		DeleteBoard(w, r)
	default:
		http.Error(w, "fuck off method not allowed", http.StatusMethodNotAllowed)

	}
}

// REMINDER: Add a captcha when everything is ready
// also, i ahve to change the way i do this
// i need to use a form request for making the site usable for everyone
// even the ones that doesnt use javascript

func PostBoard(w http.ResponseWriter, r *http.Request) {
	// decode
	var post core.Post
	if json.NewDecoder(r.Body).Decode(&post) != nil {
		http.Error(w, "the fuck is this", http.StatusBadRequest)
		return
	}
	post.Date = int(time.Now().Unix())
	// check if the board is valid
	if _, err := hex.DecodeString(post.Board); err != nil && !core.Boards[post.Board] {
		core.PrintErr("someone sent a non valid board")
		http.Error(w, "this isnt a board or a thread", http.StatusBadRequest)
		return
	}

	conns := GetRandomConns()

	blockData, _ := json.Marshal(core.BlockPost{
		Signature: crypt.SignMSG(PrivateKey, crypt.GenHashPost(post)),
		Post:      post})
	// errors
	manyErrors, status, reason := 0, 404, ""

	for _, ipConn := range conns {
		r, err := http.Post(
			fmt.Sprintf("http://%s:%d/new-post",
				ipConn.IP,
				ipConn.Port),
			"application/json", bytes.NewBuffer(blockData))
		if err != nil {
			core.PrintErr(err.Error())

			if strings.Contains(err.Error(), "connection refused") {
				delete(listConns, ipConn)
				continue
			}
			status, reason = r.StatusCode, err.Error()

			manyErrors++
		}
	}
	if manyErrors > len(conns)/2 {
		http.Error(w, reason, status)
		return
	}
	w.Write([]byte("hey thanks\n"))
}

// this is only for moderators
// this will be controlled via discord
// Im to lazy for making something better to avoid this stuff lol
func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	http.NewRequest("DELETE", "", nil)
}
