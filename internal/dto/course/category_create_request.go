package dto

import "errors"

type CategoryCreateRequest struct {
	Name string `json:"name"`
}

func (c *CategoryCreateRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name is invalid")
	}

	return nil
}
