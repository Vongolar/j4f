package gate

import (
	"net/http"
)

func listenHTTP(addr string, on func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/", on)
	http.ListenAndServe(addr, nil)
}
