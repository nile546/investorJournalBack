package models

type TableParams struct {
	Pagination Pagination  `json:"pagination"`
	Source     interface{} `json:"source"`
	SearchText string      `json:"searchText,omitempty"`
}

type Pagination struct {
	PageNumber   int `json:"pageNumber"`
	PageCount    int `json:"pageCount"`
	ItemsPerPage int `json:"itemsPerPage"`
}
