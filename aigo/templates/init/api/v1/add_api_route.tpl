package v1

import (
	"chocolate/api/v1/routers"
	"chocolate/internal/pkg/router"
)

const (
	_relativePath = "/v1"
)

type APIRouter struct{}

func (r APIRouter) Init() *router.Router {
    parent := router.NewRouterWithPath(_relativePath)
    {{range $i, $v := .TableNames}}
    parent.AddSubRouterGroup(routers.{{$v}}APIRouter{}.Init()){{end}}
    return parent
}