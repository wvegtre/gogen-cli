package user

import (
	"net/http"

	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/server/user"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/pkg/router"
)

type UsersRouter struct {
}

func (r UsersRouter) Init() *router.Router {

	h := RouterHandle{
		Service: user.NewUsersService(),
	}
	parent := router.NewRouterWithPath("/users")
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, "", h.GetUsers))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, ":id", h.GetUser))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodPost, "", h.CreateUser))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodPatch, ":id", h.UpdateUserByID))
	return parent
}
