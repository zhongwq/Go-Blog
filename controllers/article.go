package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/GoProjectGroupForEducation/Go-Blog/models"
	"github.com/GoProjectGroupForEducation/Go-Blog/services"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

func GetAllArticles(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	res, err := json.Marshal(models.GetAllArticles())
	if err != nil {
		return err
	}
	return utils.SendData(w, string(res), "OK", http.StatusOK)
}


func GetArticlesByTag(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	tag := vars["tag"]
	res, err := json.Marshal(models.GetArticlesByTag(tag))
	if err != nil {
		return err
	}
	return utils.SendData(w, string(res), "OK", http.StatusOK)
}

func CreateArticle(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	var article = models.Article{}
	author := services.GetCurrentUser(req.Header.Get("Authorization"))
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buff, &article)
	if err != nil {
		return err
	}
	article.Author = author.UserID
	id := models.CreateArticle(article)
	return utils.SendData(w, `{
  "id": `+strconv.Itoa(id)+`
}`, "OK", http.StatusOK)
}

func GetArticleByID(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	article := models.GetArticleByID(id)
	if article == nil {
		return utils.SendData(w, "{}", "Not Found", http.StatusNotFound)
	}
	author := models.GetUserByID(article.Author)
	result := make(map[string]interface{})
	result["id"] = article.ArticleID
	result["title"] = article.Title
	result["content"] = article.Content
	result["author_id"] = article.Author
	result["author"] = author.Username
	result["created_at"] = article.CreatedAt
	result["updated_at"] = article.UpdatedAt
	data, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return utils.SendData(w, string(data), "OK", http.StatusOK)
}

func UpdateArticleByID(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	var article = models.Article{}
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	article.ArticleID = id
	article.UpdatedAt = time.Now()
	err = json.Unmarshal(buff, &article)
	if err != nil {
		return err
	}
	isUpdated := models.UpdateArticleByID(id, article)
	if !isUpdated {
		id = models.CreateArticle(article)
		return utils.SendData(w, `{"id": "`+strconv.Itoa(id)+`"}`, "Created", http.StatusCreated)
	}
	return utils.SendData(w, "{}", "OK", http.StatusOK)
}
