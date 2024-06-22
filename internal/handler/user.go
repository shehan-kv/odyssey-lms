package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	queryParams "odyssey.lms/internal/dto/params"
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

	resp, err := service.GetUsers(r.Context(), queryParams.UserQueryParams{
		Page:   pageNum,
		Limit:  limitNum,
		Search: search,
		Role:   role,
	})
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.DeleteUser(r.Context(), userId)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if errors.Is(err, service.ErrLastAdminDeletion) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.ActivateUser(r.Context(), userId)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func DeactivateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.DeactivateUser(r.Context(), userId)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
