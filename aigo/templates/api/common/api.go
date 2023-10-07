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
	Page int `form:"page,omitempty" binding:"lte=100"`
	Size int `form:"size,omitempty" binding:"lte=1000"`
}

type Pagination struct {
	Page    int  `json:"page,omitempty"`
	Size    int  `json:"limit,omitempty"`
	HasNext bool `json:"has_next,omitempty"`
	Total   int  `json:"total,omitempty"`
}
