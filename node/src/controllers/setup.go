package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
	"golang.org/x/net/websocket"
)

func ConnectWS(host string) (conn *websocket.Conn, err error) {
	origin, _ := url.Parse("http://" + host)
	u, _ := url.Parse("ws://" + host)
	conn, err = websocket.DialConfig(&websocket.Config{
		Origin:    origin,
		Location:  u,
		Version:   websocket.ProtocolVersionHybi13,
		TlsConfig: nil,
	})

	return
}
func setupConns() {
	var wg sync.WaitGroup
	r, err := http.Get("http://" + core.MainServer + "/gimme5")
	if err != nil {
		log.Panic(err)
	}
	var connsIPs []core.ConnIP
	json.NewDecoder(r.Body).Decode(&connsIPs)
	Choose(connsIPs) // I update my db

	for _, c := range connsIPs {
		conn, err := ConnectWS(c.IP + ":" + strconv.Itoa(c.Port))
		if err != nil {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			Receiver(conn)
		}()

	}
	wg.Wait()
}

func setupKey() {
	r, err := http.Get("http://" + core.MainServer + "/get-info")
	if err != nil {
		log.Panic(err)
	}
	b, _ := io.ReadAll(r.Body)
	pubKey = crypt.ParsePubKey(b)
}

func Setup(port string) {
	_, err := http.Get("http://" + core.MainServer + "/connect&port=" + (port))
	if err != nil {
		panic(err)
	}
	setupKey()
	setupConns()

}
