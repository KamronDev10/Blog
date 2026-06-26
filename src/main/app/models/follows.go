package models

type Follows struct {
	FollowerId  int64 `json:"follower_id"`
	FollowingId int64 `json:"following_id"`
}
