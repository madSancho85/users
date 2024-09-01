package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"users/models"
	"users/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUsersByDateAndAge(w http.ResponseWriter, r *http.Request) {
	var dateFrom, dateTo time.Time
	var ageFrom, ageTo int
	err := json.NewDecoder(r.Body).Decode(&struct {
		DateFrom time.Time `json:"date_from"`
		DateTo   time.Time `json:"date_to"`
		AgeFrom  int       `json:"age_from"`
		AgeTo    int       `json:"age_to"`
	}{dateFrom, dateTo, ageFrom, ageTo})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users, err := h.service.GetUsersByDateAndAge(dateFrom, dateTo, ageFrom, ageTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUsersCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.GetUsersCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%d", count)
}
