package dto

import "errors"

type TicketCreateRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (t *TicketCreateRequest) Validate() error {
	if t.Subject == "" {
		return errors.New("subject is required")
	}

	if t.Description == "" {
		return errors.New("description is required")
	}

	if t.Type == "" {
		return errors.New("type is required")
	}

	if !(t.Type == "system" || t.Type == "course" || t.Type == "user") {
		return errors.New("type is invalid")
	}

	return nil
}
