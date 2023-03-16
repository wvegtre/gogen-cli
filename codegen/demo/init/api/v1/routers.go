package v1

import (
	"demo_moudle/internal/pkg/router"
)

const (
	_relativePath = "/v1"
)

type APIRouter struct{}

func (r APIRouter) Init() *router.Router {
	parent := router.NewRouterWithPath(_relativePath)
	return parent
}
