package controllers

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
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

// add it later
func GetBoard(w http.ResponseWriter, r *http.Request) {}

// REMINDER: Add a captcha when everything is ready
func PostBoard(w http.ResponseWriter, r *http.Request) {
	var post core.Post
	if json.NewDecoder(r.Body).Decode(&post) != nil {
		http.Error(w, "hijueputa enviaste una monda rara", http.StatusBadRequest)
		return
	}
	post.Date = int(time.Now().Unix())
	conns := GetRandomConns()
	// sometimes i hate and love go, in this case i hate it
	sentIt := bufio.NewReadWriter(nil, nil)

	hash := crypt.GenHashPost(post)
	signature := crypt.SignMSG(PrivateKey, hash)
	json.NewEncoder(sentIt).Encode(core.BlockPost{Signature: hex.EncodeToString(signature), Post: post})

	for _, ipConn := range conns {
		_, err := http.Post(
			fmt.Sprintf("http://%s:%d/new-post", ipConn.IP, ipConn.Port), "application/json", sentIt)
		if err != nil {
			delete(listConns, ipConn)

		}
	}
}

// this is only for moderators
// this also if you are not the problem will be reported to a discord server
// so people can know that that specific post is being reported
func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	http.NewRequest("DELETE", "", nil)
}
