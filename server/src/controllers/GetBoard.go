package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
)

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
	var blocksC [][]core.BlockPost
	for _, ipConn := range conn {
		var blocks []core.BlockPost
		r, err := http.Get(fmt.Sprintf("http://%s:%d/get-post?date=%d&board=%s", ipConn.IP, ipConn.Port, date, board))
		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				delete(listConns, ipConn)
				continue
			}
			manyErrors++
			status = r.StatusCode
			reason = err.Error()
		}
		json.NewDecoder(r.Body).Decode(&blocks)

		if !CheckValid(blocks) {
			delete(listConns, ipConn)
			continue
		}

		blocksC = append(blocksC, blocks)

	}
	if manyErrors > len(conn)/2 {
		http.Error(w, reason, status)
		return
	}
	json.NewEncoder(w).Encode(GetMostPopular(blocksC))
}

// remind me to change this to something more faster
// using go routines or something like that
func CheckValid(blocks []core.BlockPost) bool {
	for _, v := range blocks {
		signature, _ := hex.DecodeString(v.Signature)
		hash := crypt.GenHashPost(v.Post)
		if !crypt.VerifySignature(signature, hash, PublicKey) {
			return false
		}
	}
	return true
}
func GetMostPopular(blocksC [][]core.BlockPost) (final []core.BlockPost) {

	arrayMap := make(map[string]int)
	for _, array := range blocksC {

		b, _ := json.Marshal(array)
		s := string(b)
		arrayMap[s]++

	}
	s := ""
	for k, v := range arrayMap {
		if s == "" {
			s = k
		}

		if arrayMap[s] < v {
			s = k
		}
	}
	json.Unmarshal([]byte(s), &final)
	return
}
