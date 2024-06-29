package handler

import (
	"encoding/json"
	"errors"
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

func GetSupportTicketsSelf(w http.ResponseWriter, r *http.Request) {
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

	ticketRsp, err := service.GetSupportTicketsSelf(r.Context(), queryParams.TicketQueryParams{
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

func GetSupportTicketSelf(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")

	ticketId, err := strconv.ParseInt(pathId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ticketRsp, err := service.GetSupportTicketSelf(r.Context(), ticketId)
	if err != nil {
		if errors.Is(err, service.ErrNotAllowed) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ticketRsp)
}

func GetSupportTicketById(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")

	ticketId, err := strconv.ParseInt(pathId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ticketRsp, err := service.GetSupportTicketById(r.Context(), ticketId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ticketRsp)
}

func CreateSupportTicketMessage(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")

	ticketId, err := strconv.ParseInt(pathId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var createReq dto.TicketMessageCreateRequest

	err = json.NewDecoder(r.Body).Decode(&createReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createReq.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.CreateSupportTicketMessage(r.Context(), ticketId, createReq)
	if err != nil {
		if errors.Is(err, service.ErrNotAllowed) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ResolveTicket(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")

	ticketId, err := strconv.ParseInt(pathId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = service.ResolveTicket(r.Context(), ticketId)
	if err != nil {
		if errors.Is(err, service.ErrTicketNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
