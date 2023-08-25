package controllers

import (
	"encoding/json"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

func Connection(w http.ResponseWriter, r *http.Request) {
	websocket.Handler(Receiver).ServeHTTP(w, r)
}
func Sender() {
	for {
		sb := <-blocksChan
		sender, blocks := sb.Sender, sb.Blocks

		var wg sync.WaitGroup

		for c := range conns {
			wg.Add(1)
			go func(c *websocket.Conn) {
				defer wg.Done()
				if c == sender {
					return
				}
				if json.NewEncoder(c).Encode(blocks) != nil {
					delete(conns, c)
					return
				}
			}(c)

		}
		wg.Wait()
	}
}
