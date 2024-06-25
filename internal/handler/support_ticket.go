package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	queryParams "odyssey.lms/internal/dto/params"
	dto "odyssey.lms/internal/dto/ticket"
	"odyssey.lms/internal/service"
)

func CreateSupportTicket(w http.ResponseWriter, r *http.Request) {
	var createReq dto.TicketCreateRequest

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

	err = service.CreateSupportTicket(r.Context(), createReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetSupportTickets(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := query.Get("page")
	limit := query.Get("limit")
	search := query.Get("search")
	ticketType := query.Get("type")
	status := query.Get("status")

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

	ticketRsp, err := service.GetSupportTickets(r.Context(), queryParams.TicketQueryParams{
		Search: search,
		Page:   pageNum,
		Limit:  limitNum,
		Type:   ticketType,
		Status: status,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ticketRsp)
}
