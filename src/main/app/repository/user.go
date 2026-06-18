package repository

import (
	"blog_app/src/main/app/models"
	"database/sql"
)

type UserRepoI interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoI {
	return &UserRepo{db: db}
}

func (ur *UserRepo) Create(user *models.User) error {
	query := `INSERT INTO users (username , email , Password_hash) VALUES ($1 , $2 , $3 )`
	_, err := ur.db.Exec(query, user.Username, user.Email, user.Password_hash)
	if err != nil {
		return err
	}
	return nil

}

func (ur *UserRepo) GetByEmail(email string) (*models.User, error) {

	user := &models.User{}

	query := `SELECT id , username , email , Password_hash  FROM users WHERE email = $1`

	err := ur.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password_hash,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
