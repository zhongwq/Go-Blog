package main

import (
	"net/http"

	"github.com/GoProjectGroupForEducation/Go-Blog/services"

	"github.com/GoProjectGroupForEducation/Go-Blog/utils"

	"github.com/gorilla/mux"
	"github.com/GoProjectGroupForEducation/Go-Blog/controllers"
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
	return rootRouter
}

func routeArticle() {
	sub := rootRouter.PathPrefix("/api/articles").Subrouter()
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
	sub.HandleFunc("/{tag:[a-z]+}", utils.HandlerCompose(
		controllers.GetArticlesByTag,
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
		controllers.GetUserByID,
	)).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}", utils.HandlerCompose(
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
	sub := rootRouter.PathPrefix("/api/articles/{article_id:[0-9]+}/comments").Subrouter()
	sub.HandleFunc("/", utils.HandlerCompose(
		services.AuthenticationGuard,
		controllers.CreateComment,
	)).Methods("POST")
	sub.HandleFunc("/", utils.HandlerCompose(
		controllers.GetAllComments,
	)).Methods("GET")
}
