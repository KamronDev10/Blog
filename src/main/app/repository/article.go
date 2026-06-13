package repository

import (
	"blog_app/src/main/app/models"
	"database/sql"
)

type ArticleRepoI interface {
	Create(article models.Article) error
	Update(article models.Article) error
	Delete(id int) error
	Get(id int) (*models.Article, error)
	GetAll() ([]*models.Article, error)
}

type ArticleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) ArticleRepoI {
	return &ArticleRepo{db: db}
}

func (ar *ArticleRepo) Create(article models.Article) error {
	query := `
        INSERT INTO articles (title, content, view_count, active, created_at, user_id)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := ar.db.Exec(query,
		article.Title,
		article.Content,
		article.ViewCount,
		article.Active,
		article.CreatedAt,
		article.UserID,
	)
	if err != nil {
		return err
	}

	return nil
}

// -----------------------------------------------------------------------------
func (ar *ArticleRepo) GetAll() ([]*models.Article, error) {
	query := `SELECT * FROM  articles`
	rows, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}

	var articlies []*models.Article
	for rows.Next() {
		var article models.Article
		err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Content,
			&article.ViewCount,
			&article.Active,
			&article.CreatedAt,
			&article.UserID,
		)
		if err != nil {
			return nil, err
		}
		articlies = append(articlies, &article)
	}
	return articlies, nil
}

// ---------------------------------------------------------------------------------------

func (ar *ArticleRepo) Update(article models.Article) error {

	query := `UPDATE articles SET title=$1, content=$2, active=$3 WHERE id = $4`

	_, err := ar.db.Exec(
		query,
		article.Title,
		article.Content,
		article.Active,
		article.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ar *ArticleRepo) Delete(id int) error {

	query := `DELETE FROM articles WHERE id = $1`

	_, err := ar.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *ArticleRepo) Get(id int) (*models.Article, error) {
	return nil, nil
}

//   id         SERIAL PRIMARY KEY,
//     title      VARCHAR(255) NOT NULL,
//     content    TEXT NOT NULL,
//     view_count INT DEFAULT 0,
//     active     BOOLEAN DEFAULT true,
//     created_at TIMESTAMP DEFAULT NOW(),
//     user_id    INT NOT NULL,
