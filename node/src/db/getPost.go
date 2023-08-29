package db

import (
	"github.com/ranon-rat/decChan/core"
)

func GetPosts(name string, date int) (blocks []core.BlockPost) {
	db := ConnectDB()
	defer db.Close()

	if core.Boards[name] {
		r, _ := db.Query("SELECT * FROM Posts WHERE board=?1 and date<?2  ORDER BY date ASC LIMIT ?3", name, date, core.LimitPerBoard)

		blocks = ScanningPost(r)

	} else {
		// I use the board as a way of saying "hey this could be a board or a thread and if is a thread just get all this stuff and just that"
		r, _ := db.Query("SELECT * FROM Posts WHERE board=?1 OR hash=?1 ORDER BY date DESC ", name)

		blocks = ScanningPost(r)

	}

	return
}
