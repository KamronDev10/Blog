package api

import (
	"blog_app/src/main/app/api/handler"
	"net/http"
)

func Api(router *http.ServeMux, handler *handler.Handler) {
	router.HandleFunc("POST /articles/create", handler.Create)
	router.HandleFunc("GET /articles", handler.GetAll)

}
