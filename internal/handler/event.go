package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	queryParams "odyssey.lms/internal/dto/params"
	"odyssey.lms/internal/service"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := query.Get("page")
	limit := query.Get("limit")
	search := query.Get("search")
	evntType := query.Get("type")
	severity := query.Get("severity")

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

	eventRsp, err := service.GetEvents(r.Context(), queryParams.EventQueryParams{
		Page:     pageNum,
		Limit:    limitNum,
		Search:   search,
		Type:     evntType,
		Severity: severity,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&eventRsp)
}
