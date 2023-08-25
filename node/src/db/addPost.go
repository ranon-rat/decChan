package db

import (
	"github.com/ranon-rat/decChan/core"
	"github.com/ranon-rat/decChan/crypt"
)

/*
CREATE TABLE POSTS(

		Date INT,
		Body TEXT(1500),
		Username VARCHAR(64),
	    Title VARCHAR(64),
		hash varchar(64),
	    subBoard VARCHAR(64),-- it could be a response to a post, so this is something that i need to keep in mind
	    Signature VARCHAR(512)

);

CREATE TABLE Deletion(

	DatePost INT,
	DatDeletion INT,
	HashPost VARCHAR(64),
	Signature VARCHAR(512),

);
*/
func AddPost(blockPost core.BlockPost) {
	db := ConnectDB()
	post := blockPost.Post
	db.Exec("INSERT INTO POSTS(date,body,username,signature,hash,subBoard) VALUES(?1,?2,?3,?4,?5,?6,?7)", post.Date,
		post.Post,
		post.User,
		post.Title,
		blockPost.Signature,
		crypt.GenHashPost(post),
		post.SubBoard,
	)
}
func DeletePost(blockDeletion core.BlockDeletion) {
	db := ConnectDB()

	db.Exec("INSERT INTO Deletion(DatePost,DateDeletion,hashPost,signature) VALUES(?1,?2,?3,?4)", blockDeletion.DatePost, blockDeletion.DateDeletion, blockDeletion.HashPost, blockDeletion.Signature)
	db.Exec("DELETE FROM posts WHERE hash=?1 OR subBoard=?1", blockDeletion.HashPost) // everything is deleted
	// very importante importante importante :3
}
