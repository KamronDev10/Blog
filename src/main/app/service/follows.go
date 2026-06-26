package service

import (
	"blog_app/src/main/app/models"
	"blog_app/src/main/app/repository"
)

type FollowsServiceI interface {
	Follow(follow models.Follows) error
	Unfollow(followerId, followingId int64) error
	GetFollowers(userId int64) ([]*models.Follows, error)
	GetFollowing(userId int64) ([]*models.Follows, error)
}

type FollowsService struct {
	followRepo repository.FollowsRepoI
}

func NewFollowsRepo(followrepo repository.FollowsRepoI) FollowsServiceI {
	return &FollowsService{followRepo: followrepo}
}

func (fs *FollowsService) Follow(follow models.Follows) error {
	err := fs.followRepo.Follow(follow)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FollowsService) Unfollow(followerId, followingId int64) error {
	err := fs.followRepo.UnFollow(followerId, followingId)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FollowsService) GetFollowers(userId int64) ([]*models.Follows, error) {
	follows, err := fs.followRepo.GetFollowers(userId)
	if err != nil {
		return nil, err
	}
	return follows, nil
}

func (fs *FollowsService) GetFollowing(userId int64) ([]*models.Follows, error) {
	follws, err := fs.followRepo.GetFollowing(userId)
	if err != nil {
		return nil, err
	}

	return follws, nil

}
