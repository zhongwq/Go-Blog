package utils

import (
	"log"
	"net/http"
	"time"
)

func LogRequest(w http.ResponseWriter, req *http.Request, next NextFunc) error {
	var start time.Time
	var elapsed time.Duration
	defer func() {
		if err := recover(); err != nil {
			log.Println(req.Method, req.URL.String(), err.(Exception).Status, err.(Exception).Msg, elapsed)
			return
		}
		log.Println(req.Method, req.URL.String(), "200 OK", elapsed)
	}()
	start = time.Now()
	err := next()
	elapsed = time.Since(start)
	if err != nil {
		log.Println(req.Method, req.URL.String(), err, elapsed)
		return err
	}
	return nil
}
