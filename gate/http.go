package gate

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	"JFFun/task"
	"net/http"
	"strconv"
)

func listenHTTP(addr string, on func(authority string, task *task.Task, data []byte)) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cmdStr := r.Header.Get("Command")
		protoStr := r.Header.Get("Serialize-Mode")
		idStr := r.Header.Get("ID")
		authorization := r.Header.Get("Authorization")
		dataLengthStr := r.Header.Get("Data-Length")

		request := &httpRequest{
			writer:  w,
			request: r,
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			request.Reply(0, Jerror.Error_decodeError, nil)
			return
		}

		cmd, err := strconv.Atoi(cmdStr)
		if err != nil {
			request.Reply(id, Jerror.Error_decodeError, nil)
			return
		}

		proto, err := strconv.Atoi(protoStr)
		if err != nil {
			request.Reply(id, Jerror.Error_decodeError, nil)
			return
		}

		length, err := strconv.Atoi(dataLengthStr)
		if err != nil || length < 0 {
			request.Reply(id, Jerror.Error_decodeError, nil)
			return
		}

		b := make([]byte, length)

		n, err := r.Body.Read(b)
		if err != nil || n != length {
			request.Reply(id, Jerror.Error_decodeError, nil)
			return
		}

		task := &task.Task{
			ID:      id,
			CMD:     Jcommand.Command(cmd),
			SMode:   uint8(proto),
			Request: request,
		}
		on(authorization, task, b)
	})
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}

type httpRequest struct {
	writer  http.ResponseWriter
	request *http.Request
}

func (r *httpRequest) Reply(id int64, errCode Jerror.Error, data []byte) error {
	r.writer.Header().Add("Error", strconv.Itoa(int(errCode)))
	_, err := r.writer.Write(data)
	return err
}
