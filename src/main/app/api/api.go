package api

import (
	"blog_app/src/main/app/api/handler"
	"net/http"
)

func Api(router *http.ServeMux, handler *handler.Handler) {
	router.HandleFunc("POST /articles/create", handler.Create)
	router.HandleFunc("GET /articles", handler.GetAll)
	router.HandleFunc("PUT /articles/update", handler.Update)
	router.HandleFunc("DELETE /articles/delete", handler.Delete)
	router.HandleFunc("GET /articles/get", handler.Get)

}

func RegisterUserRoutes(router *http.ServeMux, h *handler.Handler) {
	router.HandleFunc("POST /auth/sign-up", h.CreateUser)

	router.HandleFunc("POST /auth/sign-in", h.Login)
}
