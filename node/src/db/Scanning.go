package db

import (
	"database/sql"
	"sync"

	"github.com/ranon-rat/decChan/core"
)

func ScanningPost(r *sql.Rows) (postsBlocks []core.BlockPost) {
	var wg sync.WaitGroup

	for r.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var block core.BlockPost

			r.Scan(&block.Post.Date,
				&block.Post.Post,
				&block.Post.User,
				&block.Post.Title,
				&block.Hash,
				&block.Post.Board,
				&block.Signature)
			postsBlocks = append(postsBlocks, block)
		}()
	}
	wg.Wait()
	return
}
func ScanningDeletes(r *sql.Rows) (delBlocks []core.BlockDeletion) {

	var wg sync.WaitGroup
	for r.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var blockDel core.BlockDeletion
			r.Scan(
				&blockDel.DatePost,
				&blockDel.DateDeletion,
				&blockDel.HashPost,
				&blockDel.Signature)

			delBlocks = append(delBlocks, blockDel)
		}()
	}
	wg.Wait()
	return
}
