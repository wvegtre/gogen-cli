package v1

import (
	"github.com/wvegtre/gogen-cli/output/easymoney/api/v1/user"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/pkg/router"
)

const (
	_relativePath = "/v1"
)

type APIRouter struct{}

func (r APIRouter) Init() *router.Router {
	parent := router.NewRouterWithPath(_relativePath)
	parent.AddSubRouterGroup(user.APIRouter{}.Init())
	return parent
}
