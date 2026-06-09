package main

import (
	_ "blog_app/docs"
	"blog_app/src/main/app/api"
	"blog_app/src/main/app/api/handler"
	"blog_app/src/main/app/repository"
	"blog_app/src/main/app/service"
	"blog_app/src/main/dependences/db"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Blog API
// @version         1.0
// @description     Blog loyihasi API
// @host            localhost:8080
// @BasePath        /

func main() {

	db, err := db.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := http.NewServeMux()
	log.Println("server is running .... ")

	{
		articleRepo := repository.NewArticleRepo(db)
		articleService := service.NewArticleService(articleRepo)

		handler := handler.Handler{
			Service: articleService,
		}

		api.Api(router, &handler)
	}
	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", router)

}
