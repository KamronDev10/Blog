package handler

import (
	"blog_app/src/main/app/dto"
	"blog_app/src/main/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Create godoc
// @Security BearerAuth
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

// qoshib
// yangi article qo'shish uchun handler func
// @Security BearerAuth
// @Summary     Article yangilash
// @Description Article title, content va active holatini yangilaydi
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id query int true "Article ID"
// @Param       article body dto.UpdateArticleRequest true "Article ma'lumotlari"
// @Success     200 {string} string "yangilandi"
// @Failure     400 {string} string "id noto'g'ri"
// @Failure     500 {string} string "server xatosi"
// @Router      /articles/update [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id noto'g'ri", http.StatusBadRequest)
		return
	}
	var newarticle dto.UpdateArticleRequest

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newarticle)

	if err != nil {
		log.Fatal(err)
	}
	err = h.Service.Update(models.Article{
		Id:      int64(id),
		Title:   newarticle.Title,
		Content: newarticle.Content,
		Active:  newarticle.Active,
	})
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("yangilandi"))
}

// artilce larni o'chirish uchun handler func
// @Security BearerAuth
// @Summary     Article o'chirish
// @Tags        articles
// @Param       id query int true "Article ID"
// @Success     200 {string} string "o'chirildi"
// @Failure     400 {string} string "id noto'g'ri"
// @Failure     500 {string} string "server xatosi"
// @Router      /articles/delete [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id noto'g'ri", http.StatusBadRequest)
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		http.Error(w, "O'chirishda xat", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("o'chirildi"))
}

//	get url uchun handler func
//
// @Summary  		bitta Article o'qish
// @Param           id query int true "Article ID"
// @Description 	Malumotlar bazasidan bitta maqolani qaytaradi
// @Tags 			articles
// @Produce 		json
// @Success      200 {array} models.Article "maqola  muvaffaqiyatli qaytarildi"
// @Failure      500 {object} map[string]string "Ichki server xatosi"
// @Router 		/articles/get [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id noto'g'ri ", http.StatusBadRequest)
		return
	}

	article, err := h.Service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ArticleResponse{
		Id:        article.Id,
		Title:     article.Title,
		Content:   article.Content,
		ViewCount: article.ViewCount,
		Active:    article.Active,
		CreatedAt: article.CreatedAt,
		UserID:    article.UserID,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)

}
