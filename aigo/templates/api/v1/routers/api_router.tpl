package routers

import (
	"net/http"
	"strconv"

	"chocolate/api/common"
	"chocolate/api/common/consts"
	"chocolate/api/v1/parameter"
	"chocolate/internal/app/services"
	"chocolate/internal/pkg/router"
	"chocolate/middleware/biz_ctx"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type {{.TableName}}APIRouter struct {}

func (r {{.TableName}}APIRouter) Init() *router.Router {

	h := {{.TableName}}APIHandler{
		Service: services.New{{.TableName}}Service(),
		Val:     validator.New(),
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
	ctx = biz_ctx.AppendFieldsToContext(ctx, "action", "get")

	id := c.Param("id")
	if err := h.Val.Var(id, consts.ValidatorTagRequired); err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	ctx = biz_ctx.AppendFieldsToContext(ctx, "id", id)

	idInt, _ := strconv.ParseInt(id, 10, 0)
	{{.TableNameLow}}Detail, err := h.Service.GetByID(ctx, idInt)
	if err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	common.ResponseOK(c, ctx, {{.TableNameLow}}Detail)
}

func (h {{.TableName}}APIHandler) List{{.TableName}}(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = biz_ctx.AppendFieldsToContext(ctx, "action", "list")

	args := new(parameter.List{{.TableName}}Args)
	args.SetDefaultPagination()

	if err := c.ShouldBindQuery(&args); err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	{{.TableNameLow}}Details, err := h.Service.List(ctx, args.ConvertToServiceArgs(),  args.Page, args.Size)
	if err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	common.ResponseOK(c, ctx, struct {
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
	ctx = biz_ctx.AppendFieldsToContext(ctx, "action", "create")

	var createArgs services.Create{{.TableName}}Args
	if err := c.ShouldBindJSON(&createArgs); err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	created{{.TableName}}, err := h.Service.Create(ctx, createArgs)
	if err != nil {
		common.ResponseError(c, ctx, err)
		return
	}

	common.ResponseCreated(c, ctx, created{{.TableName}})
}

func (h {{.TableName}}APIHandler) Update{{.TableName}}ByID(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = biz_ctx.AppendFieldsToContext(ctx, "action", "update")

	id := c.Param("id")
	if err := h.Val.Var(id, consts.ValidatorTagRequired); err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	ctx = biz_ctx.AppendFieldsToContext(ctx, "id", id)

	idInt, _ := strconv.ParseInt(id, 10, 0)

	var updateArgs services.Update{{.TableName}}Args
	if err := c.ShouldBindJSON(&updateArgs); err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	updated{{.TableName}}, err := h.Service.UpdateByID(ctx, idInt, updateArgs)
	if err != nil {
		common.ResponseError(c, ctx, err)
		return
	}

	common.ResponseOK(c, ctx, updated{{.TableName}})
}

func (h {{.TableName}}APIHandler) Delete{{.TableName}}ByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if err := h.Val.Var(id, consts.ValidatorTagRequired); err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	idInt, _ := strconv.ParseInt(id, 10, 0)
	deleted{{.TableName}}, err := h.Service.DeleteByID(ctx, idInt)
	if err != nil {
		common.ResponseError(c, ctx, err)
		return
	}
	common.ResponseOK(c, ctx, deleted{{.TableName}})
}

