package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"Go-Blog/models"

	"Go-Blog/services"
	"Go-Blog/utils"
)

// {
// 	 "content": "hello world!"
// }
func CreateComment(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	articleID, err := strconv.Atoi(vars["article_id"])
	author := services.GetCurrentUser(req.Header.Get("Authorization"))
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	comment := models.Comment{}
	err = json.Unmarshal(buff, &comment)
	if err != nil {
		panic(err)
	}
	comment.Creator = author.UserID
	comment.ArticleID = articleID
	id := models.CreateComment(&comment)
	return utils.SendData(w, `{"id":`+strconv.Itoa(id)+` }`, "OK", http.StatusOK)
}

func GetAllComments(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	articleID, err := strconv.Atoi(vars["article_id"])
	if err != nil {
		panic(err)
	}
	comments := models.GetAllCommentsByArticleID(articleID)
	var buff []interface{}
	for _, one := range comments {
		buff = append(buff, one)
	}
	data, err := json.Marshal(buff)
	if err != nil {
		panic(err)
	}
	return utils.SendData(w, string(data), "OK", http.StatusOK)
}
