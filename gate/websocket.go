package gate

import (
	Jcommand "JFFun/data/command"
	Jtask "JFFun/task"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func (m *MGate) listenWebsocket(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", m.onAcceptWebsocket)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return server.ListenAndServe()
}

func (m *MGate) onAcceptWebsocket(w http.ResponseWriter, r *http.Request) {
	authorization := strings.TrimLeft(r.URL.String(), "/")
	if len(authorization) == 0 {
		return
	}

	accID, err := m.getAccountID(authorization)
	if err != nil {
		return
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	conn := &websocketConn{
		Conn: c,
	}

	sucChan := make(chan bool)
	cr := &acceptRequest{
		accountID:  accID,
		conn:       conn,
		resultChan: sucChan,
	}
	m.acceptChan <- cr
	if suc := <-sucChan; !suc {
		return
	}
	for {
		_, b, err := conn.ReadMessage()
		if err != nil {
			m.onConnClose <- &connCloseEvent{
				accountID: accID,
				conn:      conn,
			}
			return
		}
		id, cmd, smode, raw, err := analysisBytes(b)
		if err != nil {
			break
		}
		req := &connectRequest{
			id:      id,
			cmd:     cmd,
			connect: conn,
		}
		m.requestChan <- &request{
			cmd:       cmd,
			data:      raw,
			accountID: accID,
			Request:   req,
			smode:     smode,
		}
	}
}

type websocketConn struct {
	*websocket.Conn
}

func (conn *websocketConn) sync(id uint32, cmd Jcommand.Command, resp *Jtask.ResponseData) error {
	return conn.WriteMessage(websocket.BinaryMessage, nil)
}

func (conn *websocketConn) close() error {
	return conn.Close()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	return true
}
