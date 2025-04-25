package handlers

import (
	"encoding/json"
	"learn-go/book-borrowing-api/models"
	"net/http"

	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	h.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}
