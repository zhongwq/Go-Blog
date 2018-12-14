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
		Addr:    "0.0.0.0:3001",
		Handler: RootHandler(),
	}
	fmt.Println("listen on 3001")
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
