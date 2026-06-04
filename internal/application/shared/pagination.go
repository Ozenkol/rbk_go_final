package application_shared

type Pagination struct {
	Page  int
	Limit int
}

func NewPagination(page, limit int) *Pagination {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return &Pagination{Page: page, Limit: limit}
}