package db

import "github.com/ranon-rat/decChan/core"

func ItGotToLimit(hash string) bool {
	db := ConnectDB()
	defer db.Close()
	howMany := 0
	db.QueryRow("SELECT COUNT(*) WHERE board=?1", hash).Scan(&howMany)
	return howMany < core.LimitPerThread
}
