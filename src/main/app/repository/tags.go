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
		tag := &models.Tag{}

		err := rows.Scan(
			&tag.Id,
			&tag.Name,
			&tag.Slug,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)

	}

	return tags, nil
}

func (tr *TagRepo) GetByID(id int) (*models.Tag, error) {

	query := `SELECT id , name , slug FROM tags WHERE id = $1`
	tag := &models.Tag{}
	err := tr.db.QueryRow(query, id).Scan(
		&tag.Id,
		&tag.Name,
		&tag.Slug,
	)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (tr *TagRepo) Delete(id int) error {

	query := `DELETE FROM tags WHERE id = $1`
	_, err := tr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
