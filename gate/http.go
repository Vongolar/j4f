package gate

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	Jlog "JFFun/log"
	Jtag "JFFun/log/tag"
	Jrpc "JFFun/rpc"
	Jserialization "JFFun/serialization"
	Jtask "JFFun/task"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (m *MGate) listenHTTP(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", m.analysisHTTP)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}

func (m *MGate) analysisHTTP(w http.ResponseWriter, r *http.Request) {
	cmdStr := r.Header.Get("Command")
	cmd, err := strconv.Atoi(cmdStr)
	if err != nil {
		replyHTTP(w, r.RemoteAddr, Jserialization.JSON, Jerror.Error_badRequest, nil)
		return
	}

	modeStr := r.Header.Get("Serialize-Mode")
	smode, err := strconv.Atoi(modeStr)
	if err != nil {
		replyHTTP(w, r.RemoteAddr, Jserialization.SerializateType(smode), Jerror.Error_badRequest, nil)
		return
	}

	dataLengthStr := r.Header.Get("Content-Length")
	length, err := strconv.Atoi(dataLengthStr)
	if err != nil || length < 0 {
		replyHTTP(w, r.RemoteAddr, Jserialization.SerializateType(smode), Jerror.Error_badRequest, nil)
		return
	}

	b := make([]byte, length+1)

	n, err := r.Body.Read(b)
	if (err != nil && err != io.EOF) || n != length {
		replyHTTP(w, r.RemoteAddr, Jserialization.SerializateType(smode), Jerror.Error_badRequest, nil)
		return
	}

	authorization := r.Header.Get("Authorization")
	var accID string
	if len(authorization) > 0 {
		//携带认证消息
		accID, err = m.getAccountID(authorization)
		if err != nil {
			replyHTTP(w, r.RemoteAddr, Jserialization.SerializateType(smode), Jerror.Error_unauthorized, nil)
			return
		}
	}

	Jlog.Info(Jtag.Net, fmt.Sprintf("HTTP IP %s -> %d", r.RemoteAddr, n))

	cr := Jtask.NewChannelRequest()
	request := &request{
		cmd:       Jcommand.Command(cmd),
		smode:     Jserialization.SerializateType(smode),
		accountID: accID,
		data:      b,
		Request:   cr,
	}
	m.requestChan <- request
	resp := cr.Wait()
	if resp == nil {
		replyHTTP(w, r.RemoteAddr, request.smode, Jerror.Error_undefine, nil)
		return
	}
	replyHTTP(w, r.RemoteAddr, request.smode, resp.ErrCode, resp.Data)
}

func replyHTTP(w http.ResponseWriter, remoteAddr string, st Jserialization.SerializateType, errCode Jerror.Error, data interface{}) error {
	w.Header().Add(`Error`, strconv.Itoa(int(errCode)))
	b, err := Jrpc.Encode(st, data)
	if err != nil {
		return err
	}
	n, err := w.Write(b)
	if err == nil {
		Jlog.Info(Jtag.Net, fmt.Sprintf("HTTP IP %s <- %d", remoteAddr, n))
	}
	return err
}
