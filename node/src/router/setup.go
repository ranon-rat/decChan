package router

import (
	"net/http"

	"github.com/ranon-rat/decChan/node/src/controllers"
)

func Setup(port string) {
	// this is for the server
	http.HandleFunc("/give-info-copy", func(w http.ResponseWriter, r *http.Request) {}) // 127.0.0.1:80/give-info-copy?date=12341
	http.HandleFunc("/get-post", func(w http.ResponseWriter, r *http.Request) {})       // board?board="b"&page=1&id=asdf
	http.HandleFunc("/new-post", controllers.NewPost)
	http.HandleFunc("/ws", controllers.Connection)
	http.ListenAndServe(":"+port, nil)

}
