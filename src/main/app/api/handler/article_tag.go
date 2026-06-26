package handler

import (
	"blog_app/src/main/app/models"
	"encoding/json"

	"net/http"
	"strconv"
)

// @Summary Articlega tag qo'shish
// @Tags ArticleTags
// @Accept json
// @Produce json
// @Param article_tag body dto.CreateArticleTagRequest true "Article va Tag ID"
// @Success 201 {string} string "Tag qo'shildi"
// @Security BearerAuth
// @Router /article-tags/add [post]
func (h *Handler) AddTag(w http.ResponseWriter, r *http.Request) {

	var articleTag models.ArtilceTag
	if err := json.NewDecoder(r.Body).Decode(&articleTag); err != nil {
		http.Error(w, "Noto'g'ri ma'lumot", http.StatusBadRequest)
		return
	}

	if err := h.ServiceArticleTag.AddTag(&articleTag); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Tag qo'shildi"))
}

// @Summary Article taglarini ko'rish
// @Tags ArticleTags
// @Produce json
// @Param article_id query int true "Article ID"
// @Success 200 {array} models.ArtilceTag
// @Router /article-tags [get]
func (h *Handler) GetTagsByArticleID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("article_id")
	articleId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	articleTags, err := h.ServiceArticleTag.GetByArticleID(articleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articleTags)
}

// @Summary Articledan tag o'chirish
// @Tags ArticleTags
// @Param article_id query int true "Article ID"
// @Param tag_id query int true "Tag ID"
// @Success 200 {string} string "Tag o'chirildi"
// @Security BearerAuth
// @Router /article-tags/delete [delete]
func (h *Handler) DeleteArticleTag(w http.ResponseWriter, r *http.Request) {

	articleIdStr := r.URL.Query().Get("article_id")
	tagIdStr := r.URL.Query().Get("tag_id")

	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Article ID noto'g'ri", http.StatusBadRequest)
		return
	}

	tagId, err := strconv.ParseInt(tagIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Tag ID noto'g'ri", http.StatusBadRequest)
		return
	}

	if err := h.ServiceArticleTag.DeleteTag(articleId, tagId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tag o'chirildi"))
}
