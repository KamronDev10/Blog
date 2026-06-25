package service

import (
	"blog_app/src/main/app/models"
	"blog_app/src/main/app/repository"
)

type CommentServiceI interface {
	CreateComment(comment *models.Comments) error
	GetByArticleID(articleId int64) ([]*models.Comments, error)
	Delete(id int64) error
}

type CommentService struct {
	commentrepo repository.CommentRepoI
}

func NewCommentService(commentRepo repository.CommentRepoI) CommentServiceI {
	return &CommentService{commentrepo: commentRepo}
}

func (cs *CommentService) CreateComment(comment *models.Comments) error {

}
