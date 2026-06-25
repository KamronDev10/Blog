package repository

import (
	"blog_app/src/main/app/models"
	"database/sql"
)

type CommentRepoI interface {
	CreateComment(comment *models.Comments) error
	GetByArtilceId(artilceID int64) ([]*models.Comments, error)
	Delete(id int64) error
}

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) CommentRepoI {
	return &CommentRepo{db: db}
}

func (cr *CommentRepo) CreateComment(comment *models.Comments) error {
	query := `INSERT INTO comments (content , article_id , user_id) VALUES ($1 , $2 , $3)`

	_, err := cr.db.Exec(query, comment.Content, comment.ArticleId, comment.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepo) GetByArtilceId(artilceID int64) ([]*models.Comments, error) {
	query := `SELECT id , content , article_id , user_id , created_at FROM comments WHERE article_id = $1`

	rows, err := cr.db.Query(query, artilceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comments
	for rows.Next() {
		comment := &models.Comments{}

		err := rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.ArticleId,
			&comment.UserId,
			&comment.Created_at,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (cr *CommentRepo) Delete(id int64) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := cr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
