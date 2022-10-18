package page

type Page struct {
	TotalPage  int64 `json:"total_page"`
	Page       int64 `json:"page"`
	TotalCount int64 `json:"total_count"`
	PageSize   int64 `json:"page_size"`
}
