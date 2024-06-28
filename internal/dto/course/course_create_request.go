package dto

import "errors"

type CourseCreateRequest struct {
	Name        string
	Code        string
	Description string
	Image       string
	CategoryId  int64
	Sections    []SectionCreateRequest
}

func (c *CourseCreateRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name is invalid")
	}

	if c.Code == "" {
		return errors.New("code is invalid")
	}

	if c.Description == "" {
		return errors.New("description is invalid")
	}

	for _, s := range c.Sections {
		err := s.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
