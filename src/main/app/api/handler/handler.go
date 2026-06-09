package handler

import "blog_app/src/main/app/service"

type Handler struct {
	Service service.ArticleServiceI
}
