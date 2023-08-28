package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/ranon-rat/decChan/core"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	/// CHANGE THIS LATER, JUST USE HIS IP

	ip := r.Header.Get("X-Forwarded-For") // id ont know
	portS := r.URL.Query().Get("port")
	if ip == "" || portS == "" {
		http.Error(w, "empty fields fuck you", http.StatusBadRequest)

	}
	port, err := strconv.Atoi(portS)
	if err != nil {
		http.Error(w, "fuck you you sent soemthing weird in the port field", http.StatusBadRequest)
		return
	}
	conn := core.ConnIP{IP: ip, Port: port}

	listConns[conn] = true
	fmt.Println(listConns)

	for {
		// I will change this later
		// for now i dont think that this is a problem

		out, _ := exec.Command("ping", ip, "-c 1").Output()
		if strings.Contains(string(out), "100% packet loss") {
			delete(listConns, conn)
			return
		}
		time.Sleep(time.Minute)
	}
}
