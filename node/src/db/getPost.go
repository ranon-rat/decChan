package db

import (
	"fmt"

	"github.com/ranon-rat/decChan/core"
)

func GetPostsBoard(name string, date int) (blocks core.Blocks) {
	db := ConnectDB()
	r, _ := db.Query("SELECT * FROM Posts WHERE board=?1 and date>=?2 LIMIT ?3", name, date, core.Limit)
	for r.Next() {

		r.Scan()
		fmt.Println("fuck")
	}
	return
}
func GetPostThread(hash string) {}
