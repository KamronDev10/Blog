package dto

import "time"

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Foydalanuvchi article ni tahrirlaganda
type UpdateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Active  bool   `json:"active"`
}

// Foydalanuvchiga qaytariladigan ma'lumot
type ArticleResponse struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ViewCount int64     `json:"view_count"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int64     `json:"user_id"`
}

// Ko'p article qaytarilganda
type ArticleListResponse struct {
	Articles []ArticleResponse `json:"articles"`
	Total    int64             `json:"total"`
}
