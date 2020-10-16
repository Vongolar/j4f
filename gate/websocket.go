package gate

import (
	Jerror "JFFun/data/error"
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
		// log.Error(log.TAG_CommandReq, err)
		return
	}
	onAccept(c, r)
}

func checkOrigin(r *http.Request) bool {
	return true
}

func (acc *account) listenWebsocket(conn *websocket.Conn) {
	if conn == nil {
		return
	}
	if acc.websocketConn != nil {
		acc.websocketConn.Close()
	}
	acc.websocketConn = conn
	acc.websocketConn.SetCloseHandler(func(code int, txt string) error {
		return nil
	})
	for {
		_, _, err := acc.websocketConn.ReadMessage()
		if err != nil {
			acc.websocketConn = nil
			return
		}
		// cmd := bytesToInt(data[:4])
	}
}

type websocketResp struct {
}

func (r *websocketResp) Reply(id int64, errCode Jerror.Error, data []byte) error {
	// r.writer.Header().Add("Error", strconv.Itoa(int(errCode)))
	// _, err := r.writer.Write(data)
	return nil
}
