package routers

import (
	"net/http"
	"strconv"

	"gen-templates/api/common"
	"gen-templates/api/common/consts"
	"gen-templates/api/v1/parameter"
	"gen-templates/internal/app/services"
	"gen-templates/internal/pkg/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type {{.TableName}}APIRouter struct {}

func (r {{.TableName}}APIRouter) Init() *router.Router {

	h := {{.TableName}}APIHandler{
		Service: services.New{{.TableName}}Service(),
	}
	parent := router.NewRouterWithPath("/{{.RouterPrefix}}")
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, "", h.List{{.TableName}}))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, ":id", h.Get{{.TableName}}))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodPost, "", h.Create{{.TableName}}))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodPatch, ":id", h.Update{{.TableName}}ByID))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodDelete, ":id", h.Delete{{.TableName}}ByID))
	return parent
}

type {{.TableName}}APIHandler struct {
	Service services.{{.TableName}}Service
	Val     *validator.Validate
}

func (h {{.TableName}}APIHandler) Get{{.TableName}}(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if err := h.Val.Var(id, consts.ValidatorTagRequired); err != nil {
		common.ResponseError(c, err)
		return
	}
	idInt, _ := strconv.ParseInt(id, 10, 0)
	{{.TableNameLow}}Detail, err := h.Service.GetByID(ctx, idInt)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, {{.TableNameLow}}Detail)
}

func (h {{.TableName}}APIHandler) List{{.TableName}}(c *gin.Context) {
	ctx := c.Request.Context()

	var args parameter.List{{.TableName}}Args
	if err := c.ShouldBindQuery(&args); err != nil {
		common.ResponseError(c, err)
		return
	}
	{{.TableNameLow}}Details, err := h.Service.List(ctx, args.ConvertToServiceArgs(),  args.Page, args.Size)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, struct {
		{{.TableName}}s      []*services.{{.TableName}}Model `json:"{{.TableNameLow}}s"`
		Pagination common.Pagination     `json:"pagination"`
	}{
		{{.TableName}}s: {{.TableNameLow}}Details,
		Pagination: common.Pagination{
			Page:    args.Page,
			Size:    args.Size,
			HasNext: !(len({{.TableNameLow}}Details) < args.Size),
		},
	})
}

func (h {{.TableName}}APIHandler) Create{{.TableName}}(c *gin.Context) {
	ctx := c.Request.Context()

	var createArgs services.Create{{.TableName}}Args
	if err := c.ShouldBindJSON(&createArgs); err != nil {
		common.ResponseError(c, err)
		return
	}
	created{{.TableName}}, err := h.Service.Create(ctx, createArgs)
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	common.ResponseCreated(c, created{{.TableName}})
}

func (h {{.TableName}}APIHandler) Update{{.TableName}}ByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if err := h.Val.Var(id, consts.ValidatorTagRequired); err != nil {
		common.ResponseError(c, err)
		return
	}
	idInt, _ := strconv.ParseInt(id, 10, 0)

	var updateArgs services.Update{{.TableName}}Args
	if err := c.ShouldBindJSON(&updateArgs); err != nil {
		common.ResponseError(c, err)
		return
	}
	updated{{.TableName}}, err := h.Service.UpdateByID(ctx, idInt, updateArgs)
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	common.ResponseOK(c, updated{{.TableName}})
}

func (h {{.TableName}}APIHandler) Delete{{.TableName}}ByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if err := h.Val.Var(id, consts.ValidatorTagRequired); err != nil {
		common.ResponseError(c, err)
		return
	}
	idInt, _ := strconv.ParseInt(id, 10, 0)
	deleted{{.TableName}}, err := h.Service.DeleteByID(ctx, idInt)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, deleted{{.TableName}})
}

