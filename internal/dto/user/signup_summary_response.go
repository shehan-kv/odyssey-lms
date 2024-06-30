package dto

type SignUpSummaryResponse struct {
	AllUsers       int64        `json:"allUsers"`
	Students       int64        `json:"students"`
	Administrators int64        `json:"administrators"`
	Stats          []SignUpStat `json:"stats"`
}
