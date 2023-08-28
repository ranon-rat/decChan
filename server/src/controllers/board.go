package controllers

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

// add it later
func GetBoard(w http.ResponseWriter, r *http.Request) {
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

	if _, err := hex.DecodeString(board); !core.Boards[board] && err != nil {
		http.Error(w, "non valid field, something is wrong", http.StatusBadRequest)
	}
	fmt.Println(date)
	manyErrors := 0
	reason := ""
	status := 404
	conn := GetRandomConns()
	for _, ipConn := range conn {
		r, err := http.Get(fmt.Sprintf("http://%s:%d/get-post?date=%d&board=%s", ipConn.IP, ipConn.Port, date, board))
		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				delete(listConns, ipConn)
			}
			manyErrors++
			status = r.StatusCode
			reason = err.Error()
		}
	}
	if manyErrors > len(conn)/2 {
		http.Error(w, reason, status)
	}
}

// REMINDER: Add a captcha when everything is ready
func PostBoard(w http.ResponseWriter, r *http.Request) {
	var post core.Post
	if json.NewDecoder(r.Body).Decode(&post) != nil {
		http.Error(w, "the fuck is this", http.StatusBadRequest)
		return
	}
	post.Date = int(time.Now().Unix())
	if _, err := hex.DecodeString(post.SubBoard); err != nil && !core.Boards[post.SubBoard] {
		http.Error(w, "this isnt a board or a thread this is just bullshit", http.StatusBadRequest)
		return
	}

	conns := GetRandomConns()
	// sometimes i hate and love go, in this case i hate it
	sentIt := bufio.NewReadWriter(nil, nil)

	hash := crypt.GenHashPost(post)
	signature := crypt.SignMSG(PrivateKey, hash)
	json.NewEncoder(sentIt).Encode(core.BlockPost{Signature: hex.EncodeToString(signature), Post: post})
	manyErrors := 0
	status := 404
	reason := ""
	for _, ipConn := range conns {
		r, err := http.Post(
			fmt.Sprintf("http://%s:%d/new-post", ipConn.IP, ipConn.Port), "application/json", sentIt)
		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				delete(listConns, ipConn)
			}
			status = r.StatusCode
			reason = err.Error()
			manyErrors++
		}
	}
	if manyErrors > len(conns)/2 {
		http.Error(w, reason, status)
	}
}

// this is only for moderators
// this will be controlled via discord
// Im to lazy for making something better to avoid this stuff lol
func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	http.NewRequest("DELETE", "", nil)
}
