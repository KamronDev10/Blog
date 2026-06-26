package repository

import (
	"blog_app/src/main/app/models"
	"database/sql"
)

type ArtilceTagRepoI interface {
	AddTag(artilceTag models.ArtilceTag) error
	GetByArtilceId(artilceId int64) ([]*models.ArtilceTag, error)
	DeleteTag(articleId int64, tagId int64) error
}

type ArtilceTagRepo struct {
	db *sql.DB
}

func NewArtilceTagRepo(db *sql.DB) ArtilceTagRepoI {
	return &ArtilceTagRepo{db: db}
}

func (atr *ArtilceTagRepo) AddTag(artilceTag models.ArtilceTag) error {
	query := `INSERT INTO article_tags (article_id , tag_id) VALUES ($1 , $2)`

	_, err := atr.db.Exec(query, artilceTag.ArtilceId, artilceTag.TagId)

	if err != nil {
		return err
	}

	return nil
}

func (atr *ArtilceTagRepo) GetByArtilceId(artilceId int64) ([]*models.ArtilceTag, error) {

	query := `SELECT article_id , tag_id FROM article_tags WHERE article_id = $1`

	rows, err := atr.db.Query(query, artilceId)
	if err != nil {
		return nil, err
	}

	var artilceTags []*models.ArtilceTag

	for rows.Next() {
		var articleTag models.ArtilceTag
		err := rows.Scan(
			&articleTag.ArtilceId,
			&articleTag.TagId,
		)

		if err != nil {
			return nil, err
		}
		artilceTags = append(artilceTags, &articleTag)
	}

	return artilceTags, nil
}

func (atr *ArtilceTagRepo) DeleteTag(articleId int64, tagId int64) error {

	query := `DELETE FROM article_tags WHERE article_id = $1 AND tag_id = $2`
	_, err := atr.db.Exec(query, articleId, tagId)
	if err != nil {
		return err
	}
	return nil
}
