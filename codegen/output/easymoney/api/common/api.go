package common

import "time"

type BasicResponseField struct {
	CommField
	Page *Pagination `json:"pagination"`
}

type CommField struct {
	ID       int       `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type DefaultPaginationArgs struct {
	Page  int `form:"page,omitempty" binding:"lte=100"`
	Limit int `form:"limit,omitempty" binding:"lte=1000"`
}

// 默认实现是用 has_next 来标记是否有下一页，当查询结果为空的时候，返回 false
// 实际并不准确，可能会多一次无用查询
type Pagination struct {
	Page    int  `json:"page,omitempty"`
	Limit   int  `json:"limit,omitempty"`
	HasNext bool `json:"has_next,omitempty"`
}

// 有需要的话，先执行一次 count，然后用这个 model 做 response
type PaginationWithTotal struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total int `json:"total,omitempty"`
}

func DefaultPagination() Pagination {
	return Pagination{}
}
