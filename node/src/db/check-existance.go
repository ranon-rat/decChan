package db

func CheckExistencePosts(hash string) bool {
	db := ConnectDB()
	check := 0
	db.QueryRow("SELECT 1 FROM posts WHERE hash =?1", hash).Scan(&check)
	return check > 0
}
func CheckExistenceDeletion(hash string) bool {
	db := ConnectDB()
	check := 0
	db.QueryRow("SELECT 1 FROM deletion WHERE hashPost =?1", hash).Scan(&check)
	return check > 0
}
