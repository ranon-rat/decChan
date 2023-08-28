package router

import (
	"log"
	"net/http"

	"github.com/ranon-rat/decChan/node/src/controllers"
)

func Setup(port string) {
	go controllers.Setup(port)
	go controllers.Sender()

	http.HandleFunc("/ws", controllers.Connection)
	http.HandleFunc("/give-info-copy", controllers.GiveInfo) // 127.0.0.1:80/give-info-copy?date=12341
	// this is for the server
	http.HandleFunc("/get-post", controllers.GetPosts)
	http.HandleFunc("/new-post", controllers.NewPost)
	http.HandleFunc("/del-post", controllers.DeletePost)

	log.Println(http.ListenAndServe(":"+port, nil))

}
