package dto

type CreateTagRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
