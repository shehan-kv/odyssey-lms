package dto

import "errors"

type SectionCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (s *SectionCreateRequest) Validate() error {
	if s.Title == "" {
		return errors.New("title is invalid")
	}

	return nil
}
