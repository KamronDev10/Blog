package dto

type CreateArticleTagRequest struct {
	ArticleId int64 `json:"article_id"`
	TagId     int64 `json:"tag_id"`
}
