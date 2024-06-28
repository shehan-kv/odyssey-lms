package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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

func CreateCourse(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var createReq dto.CourseCreateRequest
	createReq.Name = r.FormValue("name")
	createReq.Code = r.FormValue("code")
	createReq.Description = r.FormValue("description")
	categoryId, err := strconv.ParseInt(r.FormValue("category"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createReq.CategoryId = categoryId

	err = json.Unmarshal([]byte(r.FormValue("sections")), &createReq.Sections)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createReq.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	defer file.Close()

	cwd, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uploadsPath := filepath.Join(cwd, "uploads")

	dst, err := os.Create(filepath.Join(uploadsPath, handler.Filename))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createReq.Image = handler.Filename

	err = service.CreateCourse(r.Context(), createReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
