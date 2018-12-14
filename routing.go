package main

import (
	"net/http"
	"os"

	"github.com/GoProjectGroupForEducation/Go-Blog/services"

	"github.com/GoProjectGroupForEducation/Go-Blog/utils"

	"github.com/GoProjectGroupForEducation/Go-Blog/controllers"
	"github.com/gorilla/mux"
)

var rootRouter *mux.Router

func RootHandler() http.Handler {

	rootRouter = mux.NewRouter()
	rootRouter.HandleFunc("/help", func(w http.ResponseWriter, req *http.Request) {
		str := `API list:

GET, POST /articles 
  GET, PUT /articles/{article_id}
    GET, POST /articles/{article_id}/comments
      GET, PUT /articles/{article_id}/comments/{comment_id}


GET, POST /user
  GET, GET /user/allusers
  GET, PUT /user/{user_id}

GET, POST /tags
  GET, PUT /tags/{tag_id}

POST /auth
`
		w.WriteHeader(200)
		w.Write([]byte(str))
	})
	rootRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(utils.HandlerCompose(
			utils.LogRequest,
			utils.HandleException,
			func(w http.ResponseWriter, req *http.Request, nextFunc utils.NextFunc) error {
				next.ServeHTTP(w, req)
				return nil
			},
		))
	})
	routeArticle()
	routeComment()
	routeUser()
	routeAuth()
	routeStaticFile()
	routeDownloadFile()
	return rootRouter
}

func routeDownloadFile()  {
	//要登录才能上传文件
	sub := rootRouter.PathPrefix("/upload").Subrouter()
	sub.HandleFunc("/{filename}", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.DownloadFile,
	)).Methods("POST")
}

func routeStaticFile(){
	root, _ := os.Getwd()
	//fmt.Print(root)
	rootRouter.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(root+"/static/"))))
}

func routeArticle() {
	sub := rootRouter.PathPrefix("/articles").Subrouter()
	sub.HandleFunc("/", utils.HandlerCompose(
		controllers.GetAllArticles,
	)).Methods("GET")
	sub.HandleFunc("/", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.CreateArticle,
	)).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}", utils.HandlerCompose(
		controllers.GetArticleByID,
	)).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.UpdateArticleByID,
	)).Methods("PUT")
	sub.HandleFunc("/tag/{tag}", utils.HandlerCompose(
		controllers.GetArticlesByTag,
	)).Methods("GET")
	sub.HandleFunc("/user/{id:[0-9]+}", utils.HandlerCompose(
		controllers.GetArticlesByUserID,
	)).Methods("GET")
}

func routeUser() {
	sub := rootRouter.PathPrefix("/user").Subrouter()
	sub.HandleFunc("/allusers", utils.HandlerCompose(
		controllers.GetAllUsers,
	)).Methods("GET")
	sub.HandleFunc("/register", utils.HandlerCompose(
		controllers.CreateUser,
	)).Methods("POST")
	sub.HandleFunc("/follow", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.FollowUserByID,
	)).Methods("POST")
	sub.HandleFunc("/unfollow", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.UpdateUserByID,
	)).Methods("PUT")

	//login
	sub.HandleFunc("/login", utils.HandlerCompose(
		controllers.Auth,
	)).Methods("POST")
}

func routeAuth() {
	sub := rootRouter.PathPrefix("/api/auth").Subrouter()
	sub.HandleFunc("/", utils.HandlerCompose(
		controllers.Auth,
	)).Methods("POST")
}

func routeComment() {
	sub := rootRouter.PathPrefix("/articles/{article_id:[0-9]+}/comments").Subrouter()
	sub.HandleFunc("/", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.CreateComment,
	)).Methods("POST")
	sub.HandleFunc("/{comment_id:[0-9]+}", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.UpdateCommnetById,
	)).Methods("PUT")
	sub.HandleFunc("/", utils.HandlerCompose(
		controllers.GetAllComments,
	)).Methods("GET")
}
