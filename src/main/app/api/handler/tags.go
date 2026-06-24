package handler

import (
	"blog_app/src/main/app/dto"
	"blog_app/src/main/app/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Tag yaratish
// @Security BearerAuth
// @Tags Tags
// @Accept json
// @Produce json
// @Param tag body dto.CreateTagRequest true "Tag ma'lumotlari"
// @Success 201 {object} models.Tag
// @Failure 400 {string} string "Bad Request"
// @Security BearerAuth
// @Router /tags/create [post]
func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request) {

	var newtag dto.CreateTagRequest

	if err := json.NewDecoder(r.Body).Decode(&newtag); err != nil {
		http.Error(w, "Noto'g'ri ma'lumot", http.StatusBadRequest)
		return
	}

	if err := h.ServiceTag.CreateTag(&models.Tag{
		Name: newtag.Name,
		Slug: newtag.Slug,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newtag)
}

// @Summary Barcha taglar
// @Tags Tags
// @Produce json
// @Success 200 {array} models.Tag
// @Router /tags [get]
func (h *Handler) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.ServiceTag.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

// @Summary Bitta tag
// @Tags Tags
// @Produce json
// @Param id query int true "Tag ID"
// @Success 200 {object} models.Tag
// @Router /tags/get [get]
func (h *Handler) GetTagByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	tag, err := h.ServiceTag.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}

// @Summary Tag o'chirish
// @Security BearerAuth
// @Tags Tags
// @Param id query int true "Tag ID"
// @Success 200 {string} string "o'chirildi"
// @Security BearerAuth
// @Router /tags/delete [delete]
func (h *Handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID noto'g'ri", http.StatusBadRequest)
		return
	}

	if err := h.ServiceTag.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("o'chirildi"))
}
