package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_LogRequest(t *testing.T) {
	server := httptest.NewServer(HandlerCompose([]MiddleWare{LogRequest}))
	http.Get(server.URL)
	server.Close()
	delay := func(w http.ResponseWriter, req *http.Request, next NextFunc) error {
		time.Sleep(100000000)
		return nil
	}
	server = httptest.NewServer(HandlerCompose([]MiddleWare{LogRequest, delay}))
	http.Get(server.URL)
	server.Close()
}
