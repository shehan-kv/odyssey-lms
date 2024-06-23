package handler

import (
	"encoding/json"
	"net/http"

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
