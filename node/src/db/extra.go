package db

import "github.com/ranon-rat/decChan/core"

func ItGotToLimit(hash string) bool {
	db := ConnectDB()
	howMany := 0
	db.QueryRow("SELECT COUNT(*) WHERE board=?1", hash).Scan(&howMany)
	return howMany < core.LimitPerThread
}
