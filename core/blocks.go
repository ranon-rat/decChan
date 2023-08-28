package core

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
