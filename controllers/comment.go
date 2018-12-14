package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GoProjectGroupForEducation/Go-Blog/models"

	"github.com/GoProjectGroupForEducation/Go-Blog/services"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

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
	comment.CreatorID = author.UserID
	comment.ArticleID = articleID
	id := models.CreateComment(&comment)
	return utils.SendData(w, `{"id":`+strconv.Itoa(id)+` }`, "Create successfully", http.StatusOK)
}

func UpdateCommnetById(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	commentID, err := strconv.Atoi(vars["comment_id"])
	articleID, err := strconv.Atoi(vars["article_id"])
	author := services.GetCurrentUser(req.Header.Get("Authorization"))
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	comment := models.Comment{}
	comment.CreatorID = author.UserID
	comment.ID = commentID
	comment.ArticleID = articleID
	err = json.Unmarshal(buff, &comment)
	if err != nil {
		panic(err)
	}
	models.UpdateCommentByID(comment)
	return utils.SendData(w, "{}", "Update successfully", http.StatusOK)
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
