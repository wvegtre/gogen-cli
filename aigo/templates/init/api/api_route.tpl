package api

import (
	"net/http"

	v1 "gen-templates/init/api/v1"
	"gen-templates/internal/pkg/router"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func InitGinRouters(gdb *gorm.DB) (*gin.Engine, error) {
	e := gin.Default()
	r := routers(gdb)
	for _, subRouter := range r.SubRouterGroup {
		rg := e.Group(r.RelativePath)
		if err := includeGinRoute(rg, subRouter); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return e, nil
}

func routers(gdb *gorm.DB) *router.Router {
	r := router.NewRouterWithPath("/{{.APIPrefix}}")
	r.AddSubRouterGroup(v1.APIRouter{
		GDB: gdb,
	}.Init())
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
