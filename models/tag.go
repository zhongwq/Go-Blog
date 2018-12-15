package models

import (
	"encoding/json"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
	"strconv"
)

type Tag struct {
	Content string `json: "content"`
}

func CreateTag(tag string) {
	db := &utils.DB{}
	id := db.GenerateID("tag")
	tags := ScanTag()
	for _, val := range tags {
		if tag == val.Content {
			return
		}
	}
	db.Set("tag", strconv.Itoa(id), tag)
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