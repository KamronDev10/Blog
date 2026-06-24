package repository

import (
	"blog_app/src/main/app/models"
	"database/sql"
)

type TagRepoI interface {
	CreateTag(tag *models.Tag) error
	GetAll() ([]*models.Tag, error)
	GetByID(id int) (*models.Tag, error)
	Delete(id int) error
}

type TagRepo struct {
	db *sql.DB
}

func NewTagRepo(db *sql.DB) TagRepoI {
	return &TagRepo{db: db}
}

func (tr *TagRepo) CreateTag(tag *models.Tag) error {

	query := `SELECT INTO tags (name , slug) INTO ($1 , $2)`
	_, err := tr.db.Exec(query, tag.Name, tag.Slug)
	if err != nil {
		return err
	}

	return nil

}

func (tr *TagRepo) GetAll() ([]*models.Tag, error) {

	query := `SELECT id , name , slug FROM tags`
	rows, err := tr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*models.Tag
	for rows.Next() {
		tag := models.Tag{}

		err := rows.Scan(
			&tag.Id,
			&tag.Name,
			&tag.Slug,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)

	}

	return tags, nil
}

func (tg *TagRepo) GetByID(id int) (*models.Tag, error) {
	return nil, nil
}

func (tg *TagRepo) Delete(id int) error {
	return nil
}

// 1. CreateTag
// 2. GetAll
// 3. GetByID
// 4. Delete
