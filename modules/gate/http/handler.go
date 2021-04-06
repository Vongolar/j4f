package mhttp

import (
	"bytes"
	"io"
	"j4f/core/request"
	"j4f/core/serialize/json"
	"j4f/core/serialize/protobuf"
	"j4f/core/server"
	"j4f/core/task"
	"j4f/data/command"
	"j4f/data/errCode"
	"j4f/define"
	"net/http"
	"strconv"
)

const (
	mimeJson     = "application/json"
	mineProtobuf = "application/protobuf"
)

func (m *M_Http) GetHandlers() map[command.Command]task.TaskHandler {
	return nil
}

func (m *M_Http) getHttpMux() *http.ServeMux {
	if m.mux == nil {
		m.mux = new(http.ServeMux)
		m.mux.HandleFunc("/", m.hanldeRoute)
	}

	return m.mux
}

func (m *M_Http) hanldeRoute(w http.ResponseWriter, r *http.Request) {
	cmdStr := r.Header.Get("Command")
	cmd, err := strconv.Atoi(cmdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cr := request.CreateSyncRequest()
	t := &task.Task{CMD: command.Command(cmd), Request: cr}
	request := define.GetRequestStructByCMD(command.Command(cmd))
	if request == nil {
		b := make([]byte, r.ContentLength)
		_, err = r.Body.Read(b)
		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Data = b
	} else {
		contentType := mimeJson
		contentType = r.Header.Get("Content-Type")

		switch contentType {
		case mimeJson:
			err = json.Decode(r.Body, request)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		case mineProtobuf:
			b := make([]byte, r.ContentLength)
			_, err = r.Body.Read(b)
			if err != nil && err != io.EOF {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = protobuf.Unmarshal(b, request)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			t.Data = request
		}
	}

	server.Handle(t)
	ec, resp := cr.Wait()

	// 返回错误码
	if ec != errCode.Code_ok {
		w.Header().Set("ErrorCode", strconv.Itoa(int(ec)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 已经序列化
	if resb, ok := resp.([]byte); ok {
		//TODO:加密
		w.Write(resb)
		return
	}

	// 需要序列化
	var resb []byte
	switch r.Header.Get("Content-Type") {
	case mimeJson:
		bb := new(bytes.Buffer)
		err = json.Encode(bb, t.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resb = bb.Bytes()
	case mineProtobuf:
		resb, err = protobuf.Marshal(t.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//TODO:加密

	w.Write(resb)
}
