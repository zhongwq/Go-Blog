package controllers

import (
	"encoding/json"
	"github.com/GoProjectGroupForEducation/Go-Blog/models"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
	"net/http"
)

func GetAllTag(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	res := json.Marshal(models.GetAllTags())
	return utils.SendData(w, string(res), "OK", http.StatusOK)
}