package db

import (
	"database/sql"

	"github.com/ranon-rat/decChan/core"
)

func GetPosts(name string, date int) (blocks []core.BlockPost) {
	db := ConnectDB()
	var r *sql.Rows
	if core.Boards[name] {
		r, _ = db.Query("SELECT * FROM Posts WHERE board=?1 and date<=?2 LIMIT ?3 ORDER BY date ASC", name, date, core.LimitPerBoard)

	} else {
		// I use the board as a way of saying "hey this could be a board or a thread and if is a thread just get all this stuff and just that"
		r, _ = db.Query("SELECT * FROM Posts WHERE board=?1 OR hash=?1 ORDER BY date DESC ", name)

	}
	blocks = ScanningPost(r)

	return
}
