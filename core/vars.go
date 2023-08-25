package core

import "log"

const (
	//REMINDER, CHANGE THIS. THIS IS GAY
	MainServer  = "localhost:8080"
	DefaultPort = 8942

	ErrMsg  = "[-]"
	InfoMsg = "[?]"
	TodoMsg = "[+]"
	Limit   = 50
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
