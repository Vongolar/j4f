package jgate

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Dcommon"
	"JFFun/data/Derror"
	jschedule "JFFun/schedule"
	jserialization "JFFun/serialization"
	jtask "JFFun/task"
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
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authErr, playerID := m.authority(authorization)
	if authErr != Derror.Error_ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	conn := &websocketConn{
		Conn: c,
	}

	bindErr := m.bindConn(playerID, conn)
	if bindErr != Derror.Error_ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	for {
		_, b, err := conn.ReadMessage()

		if err != nil {
			m.unbindConn(playerID)
			return
		}

		requestData := new(Dcommon.Request)
		if err = jserialization.UnMarshal(jserialization.DefaultMode, b, requestData); err != nil {
			continue
		}

		req := &connRequest{
			id:   requestData.Id,
			conn: conn,
		}

		jschedule.HandleTaskBy(m.name, Dcommand.Command_newRequest, &jtask.Task{
			Data:     requestData,
			Request:  req,
			PlayerID: playerID,
		})
	}
}

type websocketConn struct {
	*websocket.Conn
}

func (conn *websocketConn) sync(data []byte) error {
	return conn.WriteMessage(websocket.BinaryMessage, data)
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
