package models

import (
	"encoding/json"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

type Tag struct {
	Content string `json: "content"`
}

func CreateTag(tag string) {
	db := &utils.DB{}

	if !isTagExist(tag) {
		tagItem := Tag{tag}
		buff, err := json.Marshal(tagItem)
		if err != nil {
			panic("JSON parsing error")
		}
		db.Set("tag", tag, string(buff))
	}
}

func isTagExist(content string) bool {
	db := &utils.DB{}
	buff := db.Get("tag", content)
	if len(buff) == 0 {
		return false
	}
	return true
}

func ScanTag() []Tag {
	db := &utils.DB{}
	var tagsBytes map[string]string
	var tag Tag
	var tags = []Tag{}
	tagsBytes = db.Scan("tag")
	for _, one := range tagsBytes {
		err := json.Unmarshal([]byte(one), &tag)
		tags = append(tags, tag)
		if err != nil {
			panic(err)
		}
	}
	return tags
}