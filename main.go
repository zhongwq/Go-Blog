package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	serv := &http.Server{
		Addr:    "127.0.0.1:3001",
		Handler: RootHandler(),
	}
	fmt.Println("listen on 3001")
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
