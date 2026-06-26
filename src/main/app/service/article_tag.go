package service

import (
	"blog_app/src/main/app/models"
	"blog_app/src/main/app/repository"
)

type ArtilceTagServiceI interface {
	AddTag(articleTag *models.ArtilceTag) error
	GetByArticleID(articleId int64) ([]*models.ArtilceTag, error)
	DeleteTag(articleId int64, tagId int64) error
}

type ArticleTagService struct {
	articleTagRepo repository.ArtilceTagRepoI
}

func NewArticleTagService(articleTagRepo repository.ArtilceTagRepoI) *ArticleTagService {
	return &ArticleTagService{articleTagRepo: articleTagRepo}
}

func (as *ArticleTagService) AddTag(articleTag *models.ArtilceTag) error {
	return as.articleTagRepo.AddTag(*articleTag)
}

func (as *ArticleTagService) GetByArticleID(articleId int64) ([]*models.ArtilceTag, error) {
	articleTags, err := as.articleTagRepo.GetByArtilceId(articleId)
	if err != nil {
		return nil, err
	}
	return articleTags, nil
}

func (as *ArticleTagService) DeleteTag(articleId int64, tagId int64) error {
	return as.articleTagRepo.DeleteTag(articleId, tagId)
}
