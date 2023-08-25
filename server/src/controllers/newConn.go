package controllers

import (
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
	port, err := strconv.Atoi(portS)
	if err != nil {
		http.Error(w, "fuck you you sent soemthing weird in the port field", http.StatusBadRequest)
		return
	}
	conn := core.ConnIP{IP: ip, Port: port}

	listConns[conn] = true

	for {
		// i will change this later so i can avoid using an external tool but i dont have internet right now
		// so i cant install a tool for using it
		// maybe i should do something for checking if the this is operational and its not jus a trolling node
		// sometimes it can happened but i think that the nodes will be able to handle this
		out, _ := exec.Command("ping", ip, "-c 1").Output()
		if strings.Contains(string(out), "100% packet loss") {
			delete(listConns, conn)
			return
		}
		time.Sleep(time.Minute)
	}
}
