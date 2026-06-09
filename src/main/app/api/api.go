package api

import (
	"blog_app/src/main/app/api/handler"
	"net/http"
)

func Api(router *http.ServeMux, handler *handler.Handler) {
	router.HandleFunc("/Create", handler.Create)
	router.HandleFunc("/articles", handler.GetAll)

}
