package models

type AppPaginatorInfo struct {
	HasMorePages bool  `json:"hasMorePages"`
	Page         int   `json:"page"`
	PerPage      int   `json:"perPage"`
	Total        int64 `json:"total"`
}

type AppListRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type NewAppResponse[T any] struct {
	Status        ResponseStatusType `json:"status"`
	Data          T                  `json:"data,omitempty"`
	Error         T                  `json:"error,omitempty"`
	Message       string             `json:"message,omitempty"`
	PaginatorInfo *AppPaginatorInfo  `json:"paginatorInfo,omitempty"`
}
