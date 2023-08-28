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
			var post core.Post
			var signature string
			var hash string
			r.Scan(&post.Date,
				&post.Post,
				&post.User,
				&post.Title,
				&hash,
				post.SubBoard,
				&signature)
			postsBlocks = append(postsBlocks, core.BlockPost{
				Post:      post,
				Hash:      hash,
				Signature: signature,
			})
		}()
	}
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
				blockDel.DatePost,
				blockDel.DateDeletion,
				blockDel.HashPost,
				blockDel.Signature)

			delBlocks = append(delBlocks, blockDel)
		}()
	}
	wg.Wait()
	return
}
