package handler

import (
	"encoding/json"
	"net/http"

	dto "odyssey.lms/internal/dto/course"
	"odyssey.lms/internal/service"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var createReq dto.CategoryCreateRequest

	err := json.NewDecoder(r.Body).Decode(&createReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createReq.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.CreateCategory(r.Context(), createReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := service.GetCategories(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&categories)
}
