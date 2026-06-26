package repository

import (
	"blog_app/src/main/app/models"
	"database/sql"
)

type FollowsRepoI interface {
	Follow(follow models.Follows) error
	UnFollow(followerId int64, followingId int64) error
	GetFollowers(userId int64) ([]*models.Follows, error)
	GetFollowing(userId int64) ([]*models.Follows, error)
}

type FollowRepo struct {
	db *sql.DB
}

func NewFollowerRepo(db *sql.DB) FollowsRepoI {
	return &FollowRepo{db: db}
}

func (fr *FollowRepo) Follow(follow models.Follows) error {
	query := `INSERT INTO follows (follower_id , following_id ) VALUES ($1 , $2)`
	_, err := fr.db.Exec(query, follow.FollowerId, follow.FollowingId)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FollowRepo) UnFollow(followerId int64, followingId int64) error {
	query := `DELETE FROM follows WHERE follower_id = $1 AND following_id = $2`
	_, err := fr.db.Exec(query, followerId, followingId)
	if err != nil {
		return err
	}
	return nil
}

func (fr *FollowRepo) GetFollowers(userId int64) ([]*models.Follows, error) {
	query := `SELECT follower_id , following_id FROM follows WHERE following_id = $1`
	rows, err := fr.db.Query(query, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var follows []*models.Follows

	for rows.Next() {
		var f models.Follows

		err := rows.Scan(
			&f.FollowerId,
			&f.FollowingId,
		)
		if err != nil {
			return nil, err
		}
		follows = append(follows, &f)
	}

	return follows, nil

}

func (fr *FollowRepo) GetFollowing(userId int64) ([]*models.Follows, error) {
	query := `SELECT follower_id , following_id WHERE follower_id = $1`

	rows, err := fr.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var follows []*models.Follows
	for rows.Next() {
		var f models.Follows
		err := rows.Scan(
			&f.FollowerId,
			&f.FollowingId,
		)
		if err != nil {
			return nil, err
		}
		follows = append(follows, &f)

	}

	return follows, nil

}
