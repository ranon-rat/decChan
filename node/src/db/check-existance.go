package db

func CheckExistencePosts(hash string) bool {
	db := ConnectDB()
	defer db.Close()
	check := 0
	db.QueryRow("SELECT count(*) FROM posts WHERE hash =?1", hash).Scan(&check)
	return check > 0
}
func CheckExistenceDeletion(hash string) bool {
	db := ConnectDB()
	defer db.Close()
	check := 0
	db.QueryRow("SELECT count(*) FROM deletion WHERE hashPost =?1", hash).Scan(&check)
	return check > 0
}
