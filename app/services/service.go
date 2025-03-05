package services

const (
	PageNo   = 1
	PageSize = 10
)

type Paginate struct {
	Total    int64 `json:"total"`
	PageNo   int   `json:"pageNo"`
	PageSize int   `json:"pageSize"`
	List     any   `json:"list"`
}
