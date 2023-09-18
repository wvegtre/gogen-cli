package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type Router struct {
	HttpMethod     string          `validate:"required"`
	HandleFunc     gin.HandlerFunc `validate:"required"`
	RelativePath   string          `validate:"required"`
	SubRouterGroup []*Router
}

func NewRouterWithPath(relativePatch string) *Router {
	return &Router{
		RelativePath: relativePatch,
	}
}

func NewRouter(httpMethod, relativePatch string, handle gin.HandlerFunc) *Router {
	return &Router{
		HttpMethod:   httpMethod,
		RelativePath: relativePatch,
		HandleFunc:   handle,
	}
}

func (r *Router) AddSubRouterGroup(subRouter ...*Router) {
	for _, router := range subRouter {
		r.SubRouterGroup = append(r.SubRouterGroup, router)
	}
}

func (r *Router) Validate() error {
	if err := validator.New().Struct(r); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
