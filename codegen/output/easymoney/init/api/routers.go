package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	v1 "github.com/wvegtre/gogen-cli/output/easymoney/init/api/v1"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/pkg/router"
)

func InitGinRouters() (*gin.Engine, error) {
	e := gin.Default()
	r := routers()
	for _, subRouter := range r.SubRouterGroup {
		rg := e.Group(r.RelativePath)
		if err := includeGinRoute(rg, subRouter); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return e, nil
}

func routers() *router.Router {
	r := router.NewRouterWithPath("/easy-money")
	r.AddSubRouterGroup(v1.APIRouter{}.Init())
	return r
}

func includeGinRoute(rg *gin.RouterGroup, r *router.Router) error {
	if len(r.SubRouterGroup) == 0 {
		if err := r.Validate(); err != nil {
			return errors.WithStack(err)
		}
		switch r.HttpMethod {
		case http.MethodGet:
			rg.GET(r.RelativePath, r.HandleFunc)
		case http.MethodPost:
			rg.POST(r.RelativePath, r.HandleFunc)
		case http.MethodPatch:
			rg.PATCH(r.RelativePath, r.HandleFunc)
		case http.MethodDelete:
			rg.DELETE(r.RelativePath, r.HandleFunc)
		}
		return nil
	}
	newGroup := rg.Group(r.RelativePath)
	for _, sub := range r.SubRouterGroup {
		if err := includeGinRoute(newGroup, sub); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
