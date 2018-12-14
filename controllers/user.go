package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GoProjectGroupForEducation/Go-Blog/services"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/GoProjectGroupForEducation/Go-Blog/models"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
)

func GetAllUsers(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	res, err := json.Marshal(models.GetAllUsers())
	if err != nil {
		return err
	}
	return utils.SendData(w, string(res), "OK", http.StatusOK)
}

func CreateUser(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	var user = models.User{}

	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	fmt.Println(buff);

	//json转为对应的struct
	err = json.Unmarshal(buff, &user)
	if err != nil {
		fmt.Println(err);
		return err
	}

	id := models.CreateUser(user)

	newuser := models.GetUserByID(id)
	data, err := json.Marshal(*newuser)
	if err != nil {
		return err
	}

	token := services.GenerateAuthToken(id, user.Username)
	buff, err = json.Marshal(token)

	return utils.SendData(w, `{` +
  		`"user":` + string(data) + `,` +
		`"token":` + string(buff) +
	`}`, "OK", http.StatusOK)
}

func GetUserByID(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	user := models.GetUserByID(id)
	if user == nil {
		return utils.SendData(w, "{}", "Not Found", http.StatusNotFound)
	}
	data, err := json.Marshal(*user)
	if err != nil {
		return err
	}
	return utils.SendData(w, string(data), "OK", http.StatusOK)
}


func GetUserByUsername(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	user := models.GetUserByID(id)
	if user == nil {
		return utils.SendData(w, "{}", "Not Found", http.StatusNotFound)
	}
	data, err := json.Marshal(*user)
	if err != nil {
		return err
	}
	return utils.SendData(w, string(data), "OK", http.StatusOK)
}

func UpdateUserByID(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	var user = models.User{}
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	user.UserID = id
	err = json.Unmarshal(buff, &user)
	if err != nil {
		return err
	}
	isUpdated := models.UpdateUserByID(id, user)
	//如果通过id找不到用户就创建新用户
	if !isUpdated {
		id = models.CreateUser(user)
		return utils.SendData(w, `{"id": "`+strconv.Itoa(id)+`"}`, "Created", http.StatusCreated)
	}
	return utils.SendData(w, "{}", "OK", http.StatusOK)
}


