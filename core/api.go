package core

// later i wanna improve this
// i want to make sure that this can work by connecting the nodes to the near ones
// but since this is not finish i will try to add that.
// but for now, a random connection to a node is enough
type ConnIP struct {
	IP   string
	Port int
}

type Post struct {
	User     string
	Post     string
	Date     int
	Title    string
	SubBoard string // this could be a board or a post if the post is used as a response to something
}
