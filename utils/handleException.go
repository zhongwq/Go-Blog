package utils

import (
	"log"
	"net/http"
)

func HandleException(w http.ResponseWriter, req *http.Request, next NextFunc) error {
	defer func() {
		if err := recover(); err != nil {
			SendData(w, "{}", err.(Exception).Msg, err.(Exception).Status)
			if err.(Exception).Status >= 500 {
				log.Fatal(req.Method, req.URL.String(), err.(Exception).Status, err.(Exception).Msg)
			} else {
				panic(err)
			}
		}
	}()
	err := next()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
