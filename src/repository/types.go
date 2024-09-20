package repository

type PaginationFilter struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
}

type Pagination struct {
	HasPrev  bool `json:"hasPrev"`
	PrevPage int  `json:"prevPage"`
	HasNext  bool `json:"hasNext"`
	NextPage int  `json:"nextPage"`
	CurrPage int  `json:"currPage"`
	PageSize int  `json:"pageSize"`
	LastPage int  `json:"lastPage"`
	Total    int  `json:"total"`
}

type FindOneBy struct {
	ID         uint `json:"id"`
	BusinessID uint `json:"businessId,omitempty"`
}

type FindByWithPagination struct {
	PaginationFilter
	T map[string]interface{}
}
