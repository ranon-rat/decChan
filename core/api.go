package core

type ProtocolStuff struct {
	Signature string
	BlockPost BlockPost
	DeletePost
}
type DeletePost struct {
	HashPost string
}
type BlockDeletion struct {
	HashPost string
	DatePost int

	Signature    string
	DateDeletion int
}
type BlockPost struct {
	Signature string // this will have a signature in it with rsa and a hash of the post
	Hash      string // optional
	Post      Post
}
type Blocks struct {
	BlocksPosts    []BlockPost
	BlocksDeletion []BlockDeletion
}
type ConnIP struct {
	IP   string
	Port int
}

type Gimme struct {
	Hash string
}
type Post struct {
	User     string
	Post     string
	Date     int
	Title    string
	SubBoard string // this could be a board or a post if the post is used as a response to something
}
