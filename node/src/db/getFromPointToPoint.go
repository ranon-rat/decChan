package db

import (
	"github.com/ranon-rat/decChan/core"
)

func GetPostsSince(date int) (postsBlocks []core.BlockPost) {
	db := ConnectDB()
	defer db.Close()
	r, _ := db.Query("SELECT * FROM POSTS where date>?1", date)
	// a little bit of concurrency is fine to time to time :3
	postsBlocks = ScanningPost(r)
	return
}

func GetDeleteSince(date int) (deleteBlocks []core.BlockDeletion) {
	db := ConnectDB()
	defer db.Close()
	r, _ := db.Query("SELECT * FROM Deletion where datePost>?1 and dateDeletion=?1", date)
	// a little bit of concurrency is fine to time to time :3

	deleteBlocks = ScanningDeletes(r)
	return
}
