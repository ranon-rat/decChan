package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
	"github.com/ranon-rat/decChan/node/src/db"
)

// this function let me update the database with help of the nodes
func Choose(conns []core.ConnIP) {
	lastDate := db.GetLastPostDate()

	lastDeleteDate := db.GetLastDeleteDate()
	if lastDeleteDate < lastDate {
		lastDate = lastDeleteDate
	}

	blocksPosts := make(map[core.BlockPost]bool)
	blocksDelete := make(map[core.BlockDeletion]bool)
	for _, conn := range conns {
		r, err := http.Get(fmt.Sprintf("http://%s:%d/give-info-copy?date=%d", conn.IP, conn.Port, lastDate))
		if err != nil {
			core.PrintErr("this conn has failed", conn.IP)
		}

		var blocks core.Blocks
		json.NewDecoder(r.Body).Decode(&blocks)
		for _, bp := range blocks.BlocksPosts {
			signature := hexToB(bp.Signature)
			hash := crypt.GenHashPost(bp.Post)
			if !crypt.VerifySignature(signature, hash, pubKey) {
				break
			}
			blocksPosts[bp] = true

		}
		for _, bd := range blocks.BlocksDeletion {
			signature := hexToB(bd.Signature)
			hash := crypt.GenHashDelete(bd)
			if !crypt.VerifySignature(signature, hash, pubKey) {
				break
			}
			blocksDelete[bd] = true

		}
	}
	AddDelThings(blocksPosts)
	AddDelThings(blocksDelete)

}

func AddDelThings(m any) {
	var wg sync.WaitGroup
	switch b := m.(type) {
	case map[core.BlockPost]bool:
		for bp := range b {
			wg.Add(1)

			go AddDelThing(bp, &wg)
		}
	case map[core.BlockDeletion]bool:
		for bd := range b {
			wg.Add(1)
			go AddDelThing(bd, &wg)
		}
	default:
		return
	}

	wg.Wait()
}
func AddDelThing(s any, wg *sync.WaitGroup) {
	defer wg.Done()
	switch b := s.(type) {
	case core.BlockPost:
		db.AddPost(b)
	case core.BlockDeletion:
		db.DeletePost(b)
	default:
		return
	}
}
