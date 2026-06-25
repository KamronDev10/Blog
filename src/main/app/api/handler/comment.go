package handler

import (
	"blog_app/src/main/app/dto"
	"blog_app/src/main/app/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Comment yaratish
// @Tags Comments
// @Accept json
// @Produce json
// @Param comment body dto.CreateCommentRequest true "Comment ma'lumotlari"
// @Success 201 {object} models.Comments
// @Failure 400 {string} string "Bad Request"
// @Security BearerAuth
// @Router /comments/create [post]
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Noto'g'ri ma'lumot", http.StatusBadRequest)
		return
	}

	// Middleware dan userID olish
	userID := r.Context().Value("userID").(int64)

	if err := h.ServiceComment.CreateComment(&models.Comments{
		Content:   req.Content,
		ArticleId: req.ArticleId,
		UserId:    userID,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Comment yaratildi"))
}

// @Summary Maqola commentlari
// @Tags Comments
// @Produce json
// @Param article_id query int true "Article ID"
// @Success 200 {array} models.Comments
// @Router /comments [get]
func (h *Handler) GetCommentsByArticleID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("article_id")
	articleId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	comments, err := h.ServiceComment.GetByArticleID(articleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// @Summary Comment o'chirish
// @Tags Comments
// @Param id query int true "Comment ID"
// @Success 200 {string} string "o'chirildi"
// @Security BearerAuth
// @Router /comments/delete [delete]
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	if err := h.ServiceComment.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("o'chirildi"))
}
