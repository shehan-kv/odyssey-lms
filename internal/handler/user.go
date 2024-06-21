package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	dto "odyssey.lms/internal/dto/user"
	"odyssey.lms/internal/service"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := query.Get("page")
	limit := query.Get("limit")
	search := query.Get("search")
	role := query.Get("role")

	var pageNum int
	if page == "" {
		pageNum = 1
	} else {
		num, err := strconv.Atoi(page)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pageNum = num
	}

	var limitNum int
	if limit == "" {
		limitNum = 30
	} else {
		num, err := strconv.Atoi(limit)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		limitNum = num
	}

	resp, err := service.GetUsers(r.Context(), pageNum, limitNum, search, role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var createReq dto.UserCreateRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createReq.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.CreateUser(r.Context(), createReq)
	if err != nil {
		if errors.Is(err, service.ErrInvalidRole) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
