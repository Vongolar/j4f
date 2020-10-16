package gate

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	"JFFun/serialize"
	"JFFun/task"
	"io"
	"net/http"
	"strconv"
)

func listenHTTP(addr string, on onRequestHandler) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		analysisHTTP(w, r, on)
	})

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}

//analysisHTTP 解析http请求
func analysisHTTP(w http.ResponseWriter, r *http.Request, on onRequestHandler) {
	request := &httpRequest{
		ResponseWriter: w,
	}

	cmdStr := r.Header.Get("Command")
	cmd, err := strconv.Atoi(cmdStr)
	if err != nil {
		request.Reply(&task.ResponseData{
			ErrCode: Jerror.Error_request,
		})
		return
	}

	modeStr := r.Header.Get("Serialize-Mode")
	smode, err := strconv.Atoi(modeStr)
	if err != nil {
		request.Reply(&task.ResponseData{
			ErrCode: Jerror.Error_request,
		})
		return
	}
	request.mode = smode

	authorization := r.Header.Get("Authorization")

	dataLengthStr := r.Header.Get("Data-Length")
	length, err := strconv.Atoi(dataLengthStr)
	if err != nil || length < 0 {
		request.Reply(&task.ResponseData{
			ErrCode: Jerror.Error_request,
		})
		return
	}

	b := make([]byte, length)

	n, err := r.Body.Read(b)
	if (err != nil && err != io.EOF) || n != length {
		request.Reply(&task.ResponseData{
			ErrCode: Jerror.Error_request,
		})
		return
	}
	on(authorization, request, Jcommand.Command(cmd), serialize.Mode(smode), b)
}

type httpRequest struct {
	mode serialize.Mode
	http.ResponseWriter
}

func (r *httpRequest) Reply(resp *task.ResponseData) error {
	r.Header().Add("Error", strconv.Itoa(int(resp.ErrCode)))
	serialize.Encode(r.mode, resp.Data)
	_, err := r.Write(resp.Data)
	return err
}
