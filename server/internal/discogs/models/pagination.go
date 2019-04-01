package models

type Pagination struct {
	PerPage uint `json:"per_page"`
	Items   uint `json:"items"`
	Page    uint `json:"page"`
	Pages   uint `json:"pages"`
}
