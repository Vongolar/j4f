package jgate

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Dcommon"
	"JFFun/data/Derror"
	jschedule "JFFun/schedule"
	jserialization "JFFun/serialization"
	jtask "JFFun/task"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	authorization := r.Header.Get("Authorization")
	var playerID string
	if len(authorization) > 0 {
		//有鉴权消息
		var authErr Derror.Error
		authErr, playerID = m.authority(authorization)
		if authErr != Derror.Error_ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	lengthStr := r.Header.Get("Content-Length")
	length, _ := strconv.Atoi(lengthStr)
	b := make([]byte, length)
	_, err := r.Body.Read(b)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestData := new(Dcommon.Request)
	if err = jserialization.UnMarshal(jserialization.DefaultMode, b, requestData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := &httpRequest{
		id: requestData.Id,
		r:  r,
		w:  w,
		cr: jtask.NewInnerRequest(),
	}
	jschedule.HandleTaskBy(m.name, Dcommand.Command_newRequest, &jtask.Task{
		Data:     requestData,
		Request:  request,
		PlayerID: playerID,
	})
	request.reply(request.cr.Wait())
	return
}

type httpRequest struct {
	id int32
	r  *http.Request
	w  http.ResponseWriter
	cr *jtask.InnerRequest
}

func (req *httpRequest) Reply(err Derror.Error, data interface{}) error {
	return req.cr.Reply(err, data)
}

func (req *httpRequest) reply(err Derror.Error, data interface{}) error {
	d, e := jserialization.Marshal(jserialization.DefaultMode, data)
	if e != nil {
		req.w.WriteHeader(http.StatusInternalServerError)
		return e
	}

	b, e := jserialization.Marshal(jserialization.DefaultMode, &Dcommon.Response{
		Id:   req.id,
		Err:  err,
		Data: d,
	})
	if e != nil {
		req.w.WriteHeader(http.StatusInternalServerError)
		return e
	}

	_, e = req.w.Write(b)
	return e
}
