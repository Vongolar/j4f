package gate

import (
	"JFFun/log"
	"net/http"

	"github.com/gorilla/websocket"
)

func listenWebsocket(addr string, onAccept func(conn *websocket.Conn, r *http.Request)) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		onWebsocketAccept(w, r, onAccept)
	})
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return server.ListenAndServe()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func onWebsocketAccept(w http.ResponseWriter, r *http.Request, onAccept func(conn *websocket.Conn, r *http.Request)) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(log.TAG_CommandReq, err)
		return
	}
	onAccept(c, r)
}

func checkOrigin(r *http.Request) bool {
	return true
}
