package models

import (
	"encoding/json"

	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

type AuthToken struct {
	// TokenID      int    `json:"id"`
	Token        string `json:"token"`
	AuthorizedID int    `json:"authorized_id"`
	ExpiredTime  string `json:"expired_time"`
}

func CreateToken(token AuthToken) error {
	db := &utils.DB{}
	buff, err := json.Marshal(token)
	if err != nil {
		panic(utils.Exception{err.Error(), 500})
	}
	db.Set("token", token.Token, string(buff))
	return nil
}

func GetToken(tokenstr string) *AuthToken {
	db := &utils.DB{}
	buff := db.Get("token", tokenstr)
	if len(buff) == 0 {
		return nil
	}
	token := AuthToken{}
	if err := json.Unmarshal(buff, &token); err != nil {
		panic(utils.Exception{err.Error(), 500})
	}
	return &token
}
