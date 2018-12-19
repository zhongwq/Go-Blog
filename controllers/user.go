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

//func GetAllUsers(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
//	res, err := json.Marshal(models.GetAllUsers())
//	if err != nil {
//		return err
//	}
//	return utils.SendData(w, string(res), "OK", http.StatusOK)
//}

func CreateUser(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	var user = models.User{}

	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	//json转为对应的struct
	err = json.Unmarshal(buff, &user)
	if err != nil {
		fmt.Println(err);
		return err
	}
	//不能重复用户名
	tempuser := models.GetUserByUsername(user.Username)
	if tempuser != nil{
		return utils.SendData(w, string(buff), "Username has been registered, retry.", http.StatusBadRequest)
	}
	user.Iconpath = "1.ico"
	id := models.CreateUser(user)

	newuser := models.GetUserByID(id)
	data, err := json.Marshal(*newuser)
	if err != nil {
		return err
	}

	token := services.GenerateAuthToken(&user)

	return utils.SendData(w, `{` +
  		`"user":` + string(data) + `,` +
		`"token":` + token.Token +
	`}`, "OK", http.StatusOK)
}

func FollowUserByID(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	buff, err := ioutil.ReadAll(req.Body)
	tempuser := models.User{}
	if err != nil {
		return err
	}
	err = json.Unmarshal(buff, &tempuser)
	if err != nil {
		return err
	}
	header := req.Header
	token := header.Get("Authorization")
	currentUser := services.GetCurrentUser(token)
	id := tempuser.UserID
	if err != nil {
		return err
	}
	//不能follow自己
	if id == currentUser.UserID {
		return utils.SendData(w, "", "Cannot follow yourself.", http.StatusBadRequest)
	}

	if models.Follow(id, currentUser.UserID) {
		return utils.SendData(w, "{}", "Fail when following, please check input", http.StatusNotFound)
	}

	return utils.SendData(w, "{}", "follow successfully", http.StatusOK)
}


func UnfollowUserByID(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	buff, err := ioutil.ReadAll(req.Body)
	tempuser := models.User{}
	if err != nil {
		return err
	}
	err = json.Unmarshal(buff, &tempuser)
	if err != nil {
		return err
	}
	header := req.Header
	token := header.Get("Authorization")
	currentUser := services.GetCurrentUser(token)
	id := tempuser.UserID
	if err != nil {
		return err
	}

	//不能unfollow自己
	if id == currentUser.UserID {
		return utils.SendData(w, "{}", "Cannot unfollow yourself.", http.StatusBadRequest)
	}

	if models.Unfollow(id, currentUser.UserID) {
		return utils.SendData(w, "{}", "Fail when unfollowing, please check input", http.StatusNotFound)
	}

	return utils.SendData(w, "{}", "Unfollow successfully", http.StatusOK)
}


func GetUserById(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	var user = &models.UserDetail{}
	user = models.GetUserDetailByID(id)
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
	var tmpUser = models.User{}
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	curUser := models.GetUserByID(id)
	err = json.Unmarshal(buff, &tmpUser)
	if err != nil {
		return err
	}
	if tmpUser.Password == "" {
		curUser.Email = tmpUser.Email
		curUser.Username = tmpUser.Username
	}

	isUpdated := models.UpdateUserByID(id, *curUser)
	//如果通过id找不到用户就创建新用户
	if !isUpdated {
		return utils.SendData(w, "{}", "Error when updating", http.StatusBadRequest)
	}
	return utils.SendData(w, "{}", "Update successfully!", http.StatusOK)
}



func UpdateIcon(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}
	user := models.GetUserByID(id)

	iconpath, _ := vars["filename"]

	user.Iconpath = iconpath


	err = json.Unmarshal(buff, &user)
	if err != nil {
		return err
	}
	isUpdated := models.UpdateUserByID(id, *user)
	//如果通过id找不到用户就创建新用户
	if !isUpdated {
		id = models.CreateUser(*user)
		return utils.SendData(w, `{"id": "`+strconv.Itoa(id)+`"}`, "Created", http.StatusCreated)
	}
	return utils.SendData(w, "{}", "OK", http.StatusOK)
}


func GetUserFollower(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}

	users := models.GetUserFollowers(id)

	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return utils.SendData(w, string(data), "OK", http.StatusOK)
}

func GetUserFollowing(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}

	users := models.GetUserFollowing(id)

	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return utils.SendData(w, string(data), "OK", http.StatusOK)
}



