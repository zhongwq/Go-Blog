package models

import (
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
	"time"
)

type Tag struct {
	Content string `json: "content"`
}

func CreateTag(tag string) int {
	if GetTagId(tag) == -1 {
		stmt, err := utils.GetConn().Prepare("insert into Tags values (?, ?, ?, ?)")
		if err != nil {
			panic("db insert prepare error")
		}
		res, err := stmt.Exec(nil, tag, time.Now(), time.Now())
		if err != nil {
			panic("db insert error")
		}
		tagid, err := res.LastInsertId()
		return int(tagid)
	}
	return -1
}

func GetTagId(content string) int {
	row, err := utils.GetConn().Query("SELECT t.id FROM Tags t WHERE t.content=?", content)
	if err != nil {
		panic("error")
	}
	tagid := -1
	for row.Next() {
		err = row.Scan(&tagid)
	}
	return tagid
}

func GetAllTags() []Tag {
	tags := []Tag{}
	row, err := utils.GetConn().Query("SELECT t.content FROM Tags t")
	if err != nil {
		panic("error")
	}
	for row.Next() {
		tag := Tag{}
		err = row.Scan(&tag.Content)
		tags = append(tags, tag)
	}
	return tags
}