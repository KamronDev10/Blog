package service

import (
	"blog_app/src/main/app/models"
	"blog_app/src/main/app/repository"
)

type TagServiceI interface {
	CreateTag(tag *models.Tag) error
	GetAll() ([]*models.Tag, error)
	GetByID(id int) (*models.Tag, error)
	Delete(id int) error
}

type TagService struct {
	tagrepo repository.TagRepoI
}

func NewTagService(tagrepo repository.TagRepoI) TagServiceI {
	return &TagService{tagrepo: tagrepo}
}

func (ts *TagService) CreateTag(tag *models.Tag) error {
	return ts.tagrepo.CreateTag(tag)
}

func (ts *TagService) GetAll() ([]*models.Tag, error) {

	tags, err := ts.tagrepo.GetAll()
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (ts *TagService) GetByID(id int) (*models.Tag, error) {
	tag, err := ts.tagrepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return tag, nil

}

func (ts *TagService) Delete(id int) error {
	return ts.tagrepo.Delete(id)
}
