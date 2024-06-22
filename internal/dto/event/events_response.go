package event

type EventsResponse struct {
	TotalCount int64           `json:"totalCount"`
	Events     []EventResponse `json:"events"`
}
