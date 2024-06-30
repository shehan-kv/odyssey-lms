package dto

type EventQueryParams struct {
	Page     int
	Limit    int
	Search   string
	Type     string
	Severity string
}
