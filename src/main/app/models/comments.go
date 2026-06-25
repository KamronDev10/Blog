package models

import "time"

type Comments struct {
	Id         int64     `json:"id"`
	Content    string    `json:"content"`
	ArticleId  int64     `json:"article_id"`
	UserId     int64     `json:"user_id"`
	Created_at time.Time `json:"create_at"`
}
