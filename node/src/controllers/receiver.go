package controllers

import (
	"encoding/hex"
	"encoding/json"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
	"github.com/ranon-rat/decChan/node/src/db"
	"golang.org/x/net/websocket"
)

/*
0: todo bien
1: algo repetido, continua
2: bloque no valido, mata conexion
*/
func AddSomething(hashBlock, signature []byte, hashPost string) int { // 0= to bien 1=algo repetido 2=to mal
	if !crypt.VerifySignature(signature, hashBlock, pubKey) {
		return 2
	}

	if db.CheckExistencePosts(hashPost) || db.CheckExistenceDeletion(hashPost) {
		return 1
	}
	return 0
}

// this may be used in the future i need to work on it
func Receiver(c *websocket.Conn) {
	conns[c] = true
	for {
		var block core.Blocks
		if json.NewEncoder(c).Encode(&block) != nil {
			delete(conns, c)
			return
		}
		if len(block.BlocksDeletion) == 0 && len(block.BlocksPosts) == 0 {
			core.PrintErr("empty block")
			delete(conns, c)
			return

		}
		if len(block.BlocksPosts) == 1 && len(block.BlocksDeletion) == 0 {
			blockPost := block.BlocksPosts[0]
			hashBlock := crypt.GenHashPost(blockPost.Post)
			hashPost := hex.EncodeToString(hashBlock)
			signature := hexToB(blockPost.Signature)

			switch AddSomething(hashBlock, signature, hashPost) {
			case 0:
				db.AddPost(blockPost)
				blocksChan <- BlockSender{Sender: c, Blocks: block}
			case 1:
				continue
			case 2:
				core.PrintInfo("someone sent a non valid block")
				delete(conns, c)
				return
			}

		}
		if len(block.BlocksPosts) == 0 && len(block.BlocksDeletion) == 1 {
			blockDeletion := block.BlocksDeletion[0]
			hashBlock := crypt.GenHashDelete(blockDeletion)
			hashPost := blockDeletion.HashPost
			signature := hexToB(blockDeletion.Signature)
			switch AddSomething(hashBlock, signature, hashPost) {
			case 0:
				db.DeletePost(blockDeletion)
				blocksChan <- BlockSender{Sender: c, Blocks: block}
			case 1:
				continue
			case 2:
				delete(conns, c)
				return

			}
		}
	}
}
