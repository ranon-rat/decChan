package core

import "log"

const (
	//REMINDER, CHANGE THIS. THIS IS GAY
	MainServer  = "localhost:8080"
	DefaultPort = 8942

	ErrMsg         = "[-]"
	InfoMsg        = "[?]"
	TodoMsg        = "[+]"
	LimitPerBoard  = 50
	LimitPerThread = 500
)

var (
	// later i will add more, for now is not a big deal
	Boards = map[string]bool{
		"/b/random":     true,
		"/g/technology": true,
		"/x/esoterism":  true,
		"/a/anime":      true,
	}
)

func PrintErr(msg ...any) {
	log.Println(ErrMsg, msg)
}

func PrintInfo(msg ...any) {
	log.Println(InfoMsg, msg)
}

func PrintTodo(msg ...any) {
	log.Println(TodoMsg, msg)
}
