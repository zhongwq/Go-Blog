package utils

import (
	"log"
	"net/http"
	"testing"
)

func Test_Mid(t *testing.T) {
	handler1 := func(w http.ResponseWriter, r *http.Request, next NextFunc) error {
		log.Println("Middlerware 1")
		err := next()
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println("Middlerware 1 complete")
		return nil
	}
	handler2 := func(w http.ResponseWriter, r *http.Request, next NextFunc) error {
		log.Println("Middlerware 2")
		err := next()
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println("Middlerware 2 complete")
		return nil
	}
	handler3 := func(w http.ResponseWriter, r *http.Request, next NextFunc) error {
		log.Println("Middlerware 3")
		err := next()
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println("Middlerware 3 complete")
		return nil
	}
	fun := HandlerCompose([]MiddleWare{
		handler1,
		handler2,
		handler3,
	})
	fun(nil, nil)
}
