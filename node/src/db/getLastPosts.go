package db

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
	SignatureVARCHAR(512),

);
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
