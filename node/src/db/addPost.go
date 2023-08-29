package db

import (
	"encoding/hex"

	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
)

func AddPost(blockPost core.BlockPost) {
	db := ConnectDB()
	post := blockPost.Post
	db.Exec(`INSERT INTO POSTS(
		date,
		body,
		username,
		title,
		hash,
		board,
		signature) VALUES(?1,?2,?3,?4,?5,?6,?7)`,
		post.Date,
		post.Post,
		post.User,
		post.Title,
		hex.EncodeToString(crypt.GenHashPost(post)),
		post.Board,
		blockPost.Signature,
	)
}
func DeletePost(blockDeletion core.BlockDeletion) {
	db := ConnectDB()

	db.Exec("INSERT INTO Deletion(DatePost,DateDeletion,hashPost,signature) VALUES(?1,?2,?3,?4)", blockDeletion.DatePost, blockDeletion.DateDeletion, blockDeletion.HashPost, blockDeletion.Signature)
	db.Exec("DELETE FROM posts WHERE hash=?1 OR subBoard=?1", blockDeletion.HashPost) // everything is deleted
	// very importante importante importante :3
}
