package gate

import (
	Jerror "JFFun/data/error"
	"net/http"
	"strconv"
)

func listenHTTP(addr string, on func(resp *httpResp)) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		on(&httpResp{
			writer:  w,
			request: r,
		})
	})
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}

type httpResp struct {
	writer  http.ResponseWriter
	request *http.Request
}

func (r *httpResp) Reply(id int64, errCode Jerror.Error, data []byte) error {
	r.writer.Header().Add("Error", strconv.Itoa(int(errCode)))
	_, err := r.writer.Write(data)
	return err
}
