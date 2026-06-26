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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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

	// -------------- Article Layer -----------------------
	{
		articleRepo := repository.NewArticleRepo(db)
		articleService := service.NewArticleService(articleRepo)

		handler := handler.Handler{
			Service: articleService,
		}

		api.Api(router, &handler)
	}
	// ----------------- USER layer ------------------------------------------------
	{
		userRepo := repository.NewUserRepo(db)
		userService := service.NewUserService(userRepo)

		userHandler := handler.Handler{
			ServiceUser: userService,
		}
		api.RegisterUserRoutes(router, &userHandler)
	}

	// ----------------------- tags layer --------------------------------
	{
		tagRepo := repository.NewTagRepo(db)
		tagService := service.NewTagService(tagRepo)

		tagHandler := handler.Handler{
			ServiceTag: tagService,
		}
		api.RegistrTagRoutes(router, &tagHandler)
	}

	// -------------------- comments layer ---------------------
	{
		commentRepo := repository.NewCommentRepo(db)
		commentService := service.NewCommentService(commentRepo)

		commentHandler := handler.Handler{
			ServiceComment: commentService,
		}
		api.RegistrCommentRoutes(router, &commentHandler)
	}

	// ---------------- Follows leyer -----------------------
	{
		followsRepo := repository.NewFollowerRepo(db)
		followsService := service.NewFollowsRepo(followsRepo)

		followsHandler := handler.Handler{
			ServiceFollow: followsService,
		}
		api.RegistrFollowsRoutes(router, &followsHandler)
	}

	// -- Server --------- layer
	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", router)

}
