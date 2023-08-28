
#type Post struct {
#	User     string
#	Post     string
#	Date     int
#	Title    string
#	SubBoard string // this could be a board or a post if the post is used as a response to something
#}
curl -X POST http://localhost:8080/board -H 'Content-Type: application/json' -d '{"user":"ranon","post":"testing this","title":"test","board":"/b/random"}'