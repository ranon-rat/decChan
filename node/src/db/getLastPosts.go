package db

/*
REMINDER TO END THIS AT HOME
*/
func GetLastPostDate() (date int) {
	db := ConnectDB()
	defer db.Close()
	db.QueryRow(`SELECT date FROM POSTS ORDER BY date DESC`).Scan(&date)
	return
}

func GetLastDeleteDate() (dateDeletion int) {
	db := ConnectDB()
	defer db.Close()
	db.QueryRow(`SELECT dateDeletion FROM POSTS ORDER BY date DESC`).Scan(&dateDeletion)

	return
}
