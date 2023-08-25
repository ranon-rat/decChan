package controllers

import (
	"crypto/rsa"
	"encoding/hex"

	"github.com/ranon-rat/decChan/core"
	"golang.org/x/net/websocket"
)

var (
	conns      = make(map[*websocket.Conn]bool)
	blocksChan = make(chan BlockSender)
	pubKey     *rsa.PublicKey
)

func hexToB(s string) (b []byte) {
	b, _ = hex.DecodeString(s)
	return
}

type BlockSender struct {
	Sender *websocket.Conn
	Blocks core.Blocks
}
