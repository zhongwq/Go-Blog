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

	token := services.GenerateAuthToken(id, user.Username)
	buff, err = json.Marshal(token)

	return utils.SendData(w, `{` +
  		`"user":` + string(data) + `,` +
		`"token":` + string(buff) +
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
	followUser := models.GetUserByID(id)
	//不能follow自己
	if followUser.UserID == currentUser.UserID {
		return utils.SendData(w, "", "Cannot follow yourself.", http.StatusBadRequest)
	}
	//不能重复follow
	for _, one := range followUser.Followers  {
		if one == currentUser.UserID{
			return utils.SendData(w, "{}", "You have followed him, do not follow again.", http.StatusBadRequest)
		}
	}
	followUser.Followers = append(followUser.Followers, currentUser.UserID)
	currentUser.Following = append(currentUser.Following, followUser.UserID)
	models.UpdateUserByID(followUser.UserID, *followUser)
	models.UpdateUserByID(currentUser.UserID, *currentUser)
	if followUser == nil {
		return utils.SendData(w, "{}", "id not Found", http.StatusNotFound)
	}
	if err != nil {
		return err
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
	unfollowUser := models.GetUserByID(id)
	if unfollowUser == nil {
		return utils.SendData(w, "{}", "id not Found", http.StatusNotFound)
	}
	//不能unfollow自己
	if unfollowUser.UserID == currentUser.UserID {
		return utils.SendData(w, "{}", "Cannot unfollow yourself.", http.StatusBadRequest)
	}
	//不能unfollow你没有follow的人
	index := -1
	for i, one := range unfollowUser.Followers  {
		if one == currentUser.UserID{
			index = i
		}
	}

	//不能unfollow你没有follow的人
	if index == -1{
		return utils.SendData(w, "{}", "You cannot unfollow a person whom you haven`t follow.", http.StatusBadRequest)
	}
	if index == len(unfollowUser.Followers) - 1 {
		unfollowUser.Followers = append(unfollowUser.Followers[:index], unfollowUser.Followers[index+1:]...)
	} else {
		if index == 0 {
			unfollowUser.Followers = []int{};
		} else {
			unfollowUser.Followers = unfollowUser.Followers[:index]
		}
	}

	for i, one := range currentUser.Following  {
		if one == currentUser.UserID{
			index = i
		}
	}
	if index == len(currentUser.Following) - 1 {
		currentUser.Following = append(currentUser.Following[:index], currentUser.Following[index+1:]...)
	} else {
		if index == 0 {
			currentUser.Following = []int{};
		} else {
			currentUser.Following = currentUser.Following[:index]
		}
	}


	models.UpdateUserByID(unfollowUser.UserID, *unfollowUser)
	models.UpdateUserByID(currentUser.UserID, *currentUser)

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



