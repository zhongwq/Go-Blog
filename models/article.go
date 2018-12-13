package models

import (
	"encoding/json"
	"strconv"
	"time"

	"Go-Blog/utils"
)

type ArticleList struct {
	ArticleID int    `json:"id"`
	Author    int    `json:"author_id"`
	Title     string `json:"title"`
	Tags	  []string  `json:"tags"`
	UpdatedAt string `json:"updated_at"`
}

type Article struct {
	ArticleID int       `json:"id, omitempty"`
	Author    int       `json:"author_id"`
	Title     string    `json:"title"`
	Tags	  []string  `json:"tags"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllArticles() []ArticleList {
	db := &utils.DB{}
	var articles []ArticleList
	var article ArticleList
	var articlesBytes map[string]string
	articlesBytes = db.Scan("article")
	if len(articlesBytes) == 0 {
		return []ArticleList{}
	}
	for _, one := range articlesBytes {
		err := json.Unmarshal([]byte(one), &article)
		articles = append(articles, article)
		if err != nil {
			panic(err)
		}
	}
	return articles
}

func CreateArticle(article Article) int {
	db := &utils.DB{}
	id := db.GenerateID("article")
	article.ArticleID = id
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	buff, err := json.Marshal(article)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Set("article", strconv.Itoa(id), string(buff))
	return id
}

func GetArticleByID(id int) *Article {
	db := &utils.DB{}
	buff := db.Get("article", strconv.Itoa(id))
	if len(buff) == 0 {
		return nil
	}
	article := Article{}
	err := json.Unmarshal(buff, &article)
	if err != nil {
		panic(err)
	}
	return &article
}

func UpdateArticleByID(id int, article Article) bool {
	db := &utils.DB{}
	buff := db.Get("article", strconv.Itoa(id))
	if len(buff) == 0 {
		return false
	}
	buff, err := json.Marshal(article)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Set("article", strconv.Itoa(id), string(buff))
	return true
}

func GetArticlesByTag(tag string) []ArticleList {

	db := &utils.DB{}
	var articles []ArticleList
	var article ArticleList
	var articlesBytes map[string]string
	articlesBytes = db.Scan("article")
	if len(articlesBytes) == 0 {
		return []ArticleList{}
	}
	for _, one := range articlesBytes {
		err := json.Unmarshal([]byte(one), &article)
		for _,v := range article.Tags {
			if v == tag {
				articles = append(articles, article)
			}
		}
		if err != nil {
			panic(err)
		}
	}
	return articles
}
