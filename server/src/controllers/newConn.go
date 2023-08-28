package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ranon-rat/decChan/core"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	/// CHANGE THIS LATER, JUST USE HIS IP

	ip := r.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip = r.RemoteAddr
	}

	if strings.Contains(ip, "[::1]") {
		ip = "127.0.0.1"
	}
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]
	}
	portS := r.URL.Query().Get("port")

	if portS == "" {
		http.Error(w, "empty field fuck you", http.StatusBadRequest)
		return
	}
	port, err := strconv.Atoi(portS)
	if err != nil {
		http.Error(w, "fuck you you sent soemthing weird in the port field", http.StatusBadRequest)
		return
	}
	conn := core.ConnIP{IP: ip, Port: port}

	listConns[conn] = true

}
