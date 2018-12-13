package main

import (
	"fmt"
	"testing"

	"Go-Blog/utils"
)

func Test_Auth(t *testing.T) {
	db := &utils.DB{}
	tokens := db.Scan("token")
	for k, v := range tokens {
		fmt.Printf("%s: %s\n", k, v)
	}
}
