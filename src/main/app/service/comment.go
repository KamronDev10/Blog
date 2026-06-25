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

	err := cs.commentrepo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil

}

func (cs *CommentService) GetByArticleID(articleId int64) ([]*models.Comments, error) {
	comments, err := cs.commentrepo.GetByArtilceId(articleId)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cs *CommentService) Delete(id int64) error {
	err := cs.commentrepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
