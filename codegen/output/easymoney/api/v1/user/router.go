package user

import (
	"net/http"

	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/server/user"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/pkg/router"
)

const (
	_relativePath = "/users"
)

type APIRouter struct {
}

func (r APIRouter) Init() *router.Router {

	h := UserHandle{
		Service: user.NewUsersService(),
	}
	parent := router.NewRouterWithPath(_relativePath)
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, "", h.GetUsers))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, ":id", h.GetUser))
	return parent
}
