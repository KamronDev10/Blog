package handler

import (
	"blog_app/src/main/app/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Follow qilish
// @Tags follows
// @Accept json
// @Produce json
// @Param following_id query int true "Following ID"
// @Success 200 {string} string "follow qilindi"
// @Failure 400 {string} string "id noto'g'ri"
// @Security BearerAuth
// @Router /follows/follow [post]
func (h *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	followerID := r.Context().Value("userID").(int64)

	followingStr := r.URL.Query().Get("following_id")
	followingID, err := strconv.ParseInt(followingStr, 10, 64)
	if err != nil {
		http.Error(w, "id noto'g'ri", http.StatusBadRequest)
		return
	}

	err = h.ServiceFollow.Follow(models.Follows{
		FollowerId:  followerID,
		FollowingId: followingID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("follow qilindi"))
}

// @Summary Unfollow qilish
// @Tags follows
// @Param following_id query int true "Following ID"
// @Success 200 {string} string "unfollow qilindi"
// @Failure 400 {string} string "id noto'g'ri"
// @Security BearerAuth
// @Router /follows/unfollow [delete]
func (h *Handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	followerID := r.Context().Value("userID").(int64)

	followingStr := r.URL.Query().Get("following_id")
	followingID, err := strconv.ParseInt(followingStr, 10, 64)
	if err != nil {
		http.Error(w, "id noto'g'ri", http.StatusBadRequest)
		return
	}

	err = h.ServiceFollow.Unfollow(followerID, followingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("unfollow qilindi"))
}

// @Summary Obunachlarni ko'rish
// @Tags follows
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {array} models.Follows
// @Router /follows/followers [get]
func (h *Handler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	userStr := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userStr, 10, 64)
	if err != nil {
		http.Error(w, "id noto'g'ri", http.StatusBadRequest)
		return
	}

	followers, err := h.ServiceFollow.GetFollowers(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(followers)
}

// @Summary Following larni ko'rish
// @Tags follows
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {array} models.Follows
// @Router /follows/following [get]
func (h *Handler) GetFollowing(w http.ResponseWriter, r *http.Request) {
	userStr := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userStr, 10, 64)
	if err != nil {
		http.Error(w, "id noto'g'ri", http.StatusBadRequest)
		return
	}

	following, err := h.ServiceFollow.GetFollowing(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(following)
}
