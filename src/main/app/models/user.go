package models

import "time"

type User struct {
	Id            int64  `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password_hash string `json:"password_hash"`
	Avatar        string `json:"avatar"`
	Bio           string `json:"bio,omitempty"`
	Registered_at time.Time
}

// CREATE TABLE users (
//     id            SERIAL PRIMARY KEY,
//     username      VARCHAR(50) UNIQUE NOT NULL,
//     email         VARCHAR(100) UNIQUE NOT NULL,
//     password      VARCHAR(255) NOT NULL,
//     avatar        VARCHAR(255),
//     bio           TEXT,
//     registered_at TIMESTAMP DEFAULT NOW()
// );
