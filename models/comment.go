package models

import (
	"encoding/json"
	"strconv"
	"time"

	"Go-Blog/utils"
)

type Comment struct {
	CommentID int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID int       `json:"article_id"`
	Creator   int       `json:"creator"`
}

func CreateComment(comment *Comment) int {
	db := &utils.DB{}
	id := db.GenerateID("comment")
	comment.CommentID = id
	comment.CreatedAt = time.Now()
	buff, err := json.Marshal(comment)
	if err != nil {
		panic(err)
	}
	db.Set("comment", strconv.Itoa(id), string(buff))
	return id
}

func GetAllCommentsByArticleID(articleID int) []Comment {
	db := &utils.DB{}
	comments := db.Scan("comment")
	var result []Comment
	for _, v := range comments {
		tmp := Comment{}
		err := json.Unmarshal([]byte(v), &tmp)
		if err != nil {
			panic(err)
		}
		if tmp.ArticleID == articleID {
			result = append(result, tmp)
		}
	}
	return result
}
