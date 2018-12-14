package services

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/GoProjectGroupForEducation/Go-Blog/utils"

	"github.com/GoProjectGroupForEducation/Go-Blog/models"
)

var checksum = []byte("heng-is-a-very-handsome-boy")

func GenerateAuthToken(userID int, username string) models.AuthToken {
	h := md5.New()
	expiredTime := time.Now().UnixNano()/1e6 + 1000*60*60*3 // expired time is 3 hours
	source := strconv.FormatInt(expiredTime, 10) + strconv.Itoa(userID)
	io.WriteString(h, source)
	token := fmt.Sprintf("%x", h.Sum(checksum))
	authToken := models.AuthToken{
		// 0,
		token,
		userID,
		username,
		strconv.FormatInt(expiredTime, 10),
	}
	models.CreateToken(authToken)
	// authToken.TokenID = id
	return authToken
}

func authenticateToken(token string) bool {
	// id := token.TokenID
	real := models.GetToken(token)
	if real == nil {
		return false
	}
	// if (token.Token != real.Token) ||
	// 	(token.AuthorizedID != real.AuthorizedID) ||
	// 	(token.ExpiredTime != real.ExpiredTime) {
	// 	return false
	// }
	if real.ExpiredTime < strconv.FormatInt((time.Now().UnixNano()/1e6), 10) {
		return false
	}

	return true
}

func GetCurrentUser(token string) *models.User {
	data := models.GetToken(token)
	user := models.GetUserByID(data.AuthorizedID)
	return user
}

func AuthenticationGuard(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	header := req.Header
	token := header.Get("Authorization")
	if authenticateToken(token) {
		return next()
	} else {
		panic(utils.Exception{"Unauthorized", http.StatusUnauthorized})
	}
}
