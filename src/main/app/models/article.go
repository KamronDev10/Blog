package models

import "time"

type Article struct {
	Id        int64
	Title     string
	Content   string
	ViewCount int64
	Active    bool
	CreatedAt time.Time
	UserID    int64
}

// id         SERIAL PRIMARY KEY,
// title      VARCHAR(255) NOT NULL,
// content    TEXT NOT NULL,
// view_count INT DEFAULT 0,
// active     BOOLEAN DEFAULT true,
// created_at TIMESTAMP DEFAULT NOW(),
// user_id    INT NOT NULL,
