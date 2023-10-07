package v1

import (
	"chocolate/api/v1/routers"
	"chocolate/internal/pkg/router"

	"gorm.io/gorm"
)

const (
	_relativePath = "/v1"
)

type APIRouter struct{
    GDB *gorm.DB
}

func (r APIRouter) Init() *router.Router {
    parent := router.NewRouterWithPath(_relativePath)
    {{range $i, $v := .TableNames}}
        parent.AddSubRouterGroup(routers.{{$v}}APIRouter{}.Init())
    {{end}}
    return parent
}