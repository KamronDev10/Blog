package handler

import (
	"blog_app/src/main/app/dto"
	"blog_app/src/main/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Create godoc
// @Summary     Yangi article yaratish
// @Description Title va content bilan yangi article yaratadi
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body dto.CreateArticleRequest true "Article ma'lumotlari"
// @Success     201 {string} string "muvafaqiyatli qo'shildi"
// @Failure     400 {string} string "Noto'g'ri so'rov"
// @Failure     500 {string} string "Server xatosi"
// @Router      /articles/create [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var newArticle dto.CreateArticleRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newArticle)
	if err != nil {
		log.Fatal(err)
	}

	err = h.Service.Create(
		&models.Article{
			Title:   newArticle.Title,
			Content: newArticle.Content,
			UserID:  1,
		},
	)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("muvafaqiyatli qo'shildi "))
}

// GetAll godoc
// @Summary      Barcha maqolalarni olish
// @Description  Ma'lumotlar bazasidan barcha maqolalar ro'yxatini qaytaradi
// @Tags         articles
// @Produce      json
// @Success      200 {array} models.Article "Maqolalar ro'yxati muvaffaqiyatli qaytarildi"
// @Failure      500 {object} map[string]string "Ichki server xatosi"
// @Router       /articles [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	articles, err := h.Service.GetAll()
	if err != nil {
		fmt.Println("err", err)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(articles)
}
