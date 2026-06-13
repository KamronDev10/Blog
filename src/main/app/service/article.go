package service

import (
	"blog_app/src/main/app/models"
	"blog_app/src/main/app/repository"
)

type ArticleServiceI interface {
	Create(article *models.Article) error
	GetAll() ([]*models.Article, error)
	Update(article models.Article) error
	Delete(id int) error
}

type articleService struct {
	articleRepo repository.ArticleRepoI
}

func NewArticleService(articleRepo repository.ArticleRepoI) ArticleServiceI {
	return &articleService{articleRepo: articleRepo}
}

func (as *articleService) Create(article *models.Article) error {
	return as.articleRepo.Create(*article)
}

func (as *articleService) GetAll() ([]*models.Article, error) {
	return as.articleRepo.GetAll()
}

func (as *articleService) Update(article models.Article) error {
	return as.articleRepo.Update(article)
}

func (as *articleService) Delete(id int) error {
	return as.articleRepo.Delete(id)
}
