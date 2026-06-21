package api

import (
	"blog_app/src/main/app/api/handler"
	"blog_app/src/main/common/middleware"
	"net/http"
)

func Api(router *http.ServeMux, handler *handler.Handler) {
	router.Handle(
		"POST /articles/create",
		middleware.AuthMiddleware(http.HandlerFunc(handler.Create)))

	router.Handle(
		"PUT /articles/update",
		middleware.AuthMiddleware(http.HandlerFunc(handler.Update)))

	router.Handle(
		"DELETE /articles/delete",
		middleware.AuthMiddleware(http.HandlerFunc(handler.Delete)))

	// -------------------------------------------------------------  not middle ware
	router.HandleFunc("GET /articles", handler.GetAll)
	router.HandleFunc("GET /articles/get", handler.Get)

}

func RegisterUserRoutes(router *http.ServeMux, h *handler.Handler) {
	router.HandleFunc("POST /auth/sign-up", h.CreateUser)

	router.HandleFunc("POST /auth/sign-in", h.Login)
}
