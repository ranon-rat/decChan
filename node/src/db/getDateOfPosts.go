package db

func GetDatePost(hash string) (date int) {
	db := ConnectDB()
	defer db.Close()
	db.QueryRow("SELECT date FROM Posts WHERE hash=?1", hash).Scan(&date)
	return
}
