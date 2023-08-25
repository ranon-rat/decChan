package db

import (
	"sync"

	"github.com/ranon-rat/decChan/core"
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
func GetPostsSince(date int) (postsBlocks []core.BlockPost) {
	db := ConnectDB()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM POSTS where date>?1", date)
	// a little bit of concurrency is fine to time to time :3
	var wg sync.WaitGroup

	for rows.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var post core.Post
			var signature string
			var hash string
			rows.Scan(&post.Date, &post.Post, &post.User, &post.Title, &hash, post.SubBoard, &signature)
			postsBlocks = append(postsBlocks, core.BlockPost{
				Post:      post,
				Hash:      hash,
				Signature: signature,
			})
		}()
	}
	wg.Wait()
	return
}

/*
CREATE TABLE Deletion(
    DatePost int,
    DateDeletion INT,
    HashPost VARCHAR(64) unique,
    Signature VARCHAR(512) UNIQUE

);
*/

func GetDeleteSince(date int) (deleteBlocks []core.BlockDeletion) {
	db := ConnectDB()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM Deletion where datePost>?1 and dateDeletion=?1", date)
	// a little bit of concurrency is fine to time to time :3
	var wg sync.WaitGroup

	for rows.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var blockDelete core.BlockDeletion
			rows.Scan(blockDelete.DatePost, blockDelete.DateDeletion, blockDelete.HashPost, blockDelete.Signature)

			deleteBlocks = append(deleteBlocks, blockDelete)
		}()
	}
	wg.Wait()
	return
}
