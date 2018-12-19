package models

import (
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

type Tag struct {
	Content string `json: "content"`
}

func CreateTag(tag string) int {
	stmt, err := utils.GetConn().Prepare("insert into Tags values (?)")
	if err != nil {
		panic("db insert prepare error")
	}
	res, err := stmt.Exec(tag)
	if err != nil {
		panic("db insert error")
	}
	tagid, err := res.LastInsertId()
	return int(tagid)
}


func GetAllTags() []Tag {
	tags := []Tag{}
	row, err := utils.GetConn().Query("SELECT content FROM Tags")
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