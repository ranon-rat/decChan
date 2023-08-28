package router

import (
	"net/http"

	"github.com/ranon-rat/decChan/server/src/controllers"
)

func Setup(port string) {
	controllers.SetupRSA() // this is really importante
	http.HandleFunc("/get-info", controllers.GetInfo)
	http.HandleFunc("/gimme5", controllers.Conns5)
	http.HandleFunc("/connect", controllers.Connect) // this will work for connecting and doing gods work
	http.HandleFunc("/board", controllers.Board)     // /board?name="B"&id="ab123dda123fff1"

	http.ListenAndServe(":"+port, nil)
}
