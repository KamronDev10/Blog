package dto

type CreateCommentRequest struct {
	Content   string `json:"content"`
	ArticleId int64  `json:"article_id"`
}
