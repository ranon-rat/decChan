CREATE TABLE POSTS(

	Date INT,
	Body TEXT(1500),
	Username VARCHAR(64),
    Title VARCHAR(64),
	hash varchar(64) UNIQUE,
    Board VARCHAR(64),-- it could be a response to a post, so this is something that i need to keep in mind
    Signature VARCHAR(512) unique
    

);

CREATE TABLE Deletion(
    DatePost int,
    DateDeletion INT,
    HashPost VARCHAR(64) unique,
    Signature VARCHAR(512) UNIQUE 

);