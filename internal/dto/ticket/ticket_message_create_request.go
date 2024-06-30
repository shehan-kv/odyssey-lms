package dto

import "errors"

type TicketMessageCreateRequest struct {
	Message string `json:"message"`
}

func (t *TicketMessageCreateRequest) Validate() error {
	if t.Message == "" {
		return errors.New("message in invalid")
	}

	return nil
}
