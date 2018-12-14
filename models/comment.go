package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

type CommentList struct {
	ID 		  int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID int       `json:"article_id"`
	Creator   UserList  `json:"creator"`
	CreatorID int       `json:"creator_id"`
}

type Comment struct {
	ID 		  int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID int       `json:"article_id"`
	CreatorID int       `json:"creator_id"`
}

func CreateComment(comment *Comment) int {
	db := &utils.DB{}
	id := db.GenerateID("comment")
	comment.ID = id
	comment.CreatedAt = time.Now()
	buff, err := json.Marshal(comment)
	if err != nil {
		panic(err)
	}
	db.Set("comment", strconv.Itoa(id), string(buff))
	return id
}

func GetAllCommentsByArticleID(articleID int) []CommentList {
	db := &utils.DB{}
	comments := db.Scan("comment")
	var result = []CommentList{}
	for _, v := range comments {
		tmp := CommentList{}
		err := json.Unmarshal([]byte(v), &tmp)
		tmp.Creator = *GetUserListByID(tmp.CreatorID)
		if err != nil {
			panic(err)
		}
		if tmp.ArticleID == articleID {
			result = append(result, tmp)
		}
	}
	return result
}

func UpdateCommentByID(comment Comment) bool {
	db := &utils.DB{}
	buff := db.Get("comment", strconv.Itoa(comment.ID))
	if len(buff) == 0 {
		return false
	}
	buff, err := json.Marshal(comment)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Set("comment", strconv.Itoa(comment.ID), string(buff))
	return true
}
