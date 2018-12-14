package controllers

import (
	"encoding/json"
	"github.com/GoProjectGroupForEducation/Go-Blog/services"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

func Auth(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(buff, &body)

	if err != nil {
		return err
	}

	if _, ok := body["id"]; ok {
		//存在
		id := int(body["id"].(float64))
		token := services.GenerateAuthToken(id)
		//json 转为定义的token类
		buff, err = json.Marshal(token)
		return utils.SendData(w, string(buff), "OK", http.StatusOK)
	}else {
		log.Println(req.Method, req.URL.String(), http.StatusBadRequest)
		return utils.SendData(w, string(buff), "Id not found", http.StatusBadRequest)
	}
}
