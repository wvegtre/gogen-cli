{{define "convert_args"}}
type RouterHandle struct {
	Service *{{.group}}.{{.service_prefix}}Service
	Val     *validator.Validate
}

func (h RouterHandle) Get{{.resource}}ByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := common.ParseAndValidateParamID(c.Param("id"), h.Val)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	{{.detail_prefix}}Detail, err := h.Service.Get(ctx, id)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, {{.detail_prefix}}Detail)
}

func (h RouterHandle) Get{{.resource}}s(c *gin.Context) {
	ctx := c.Request.Context()

	var args Get{{.resource}}sArgs
	if err := c.ShouldBindQuery(args); err != nil {
		common.ResponseError(c, err)
		return
	}
	{{.detail_prefix}}Details, err := h.Service.List(ctx, args.ConvertToServiceArgs(), args.GenQueryOptions()...)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, {{.resource}}sListResponse{
		{{.resource}}s:      {{.detail_prefix}}Details,
		Pagination: common.Pagination{},
	})
}

func (h RouterHandle) Create{{.resource}}(c *gin.Context) {
	ctx := c.Request.Context()

	var args {{.group}}.Create{{.resource}}sArgs
	if err := c.ShouldBindQuery(args); err != nil {
		common.ResponseError(c, err)
		return
	}
	{{.detail_prefix}}Details, err := h.Service.Create(ctx, args)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, {{.detail_prefix}}Details)
}

func (h RouterHandle) Update{{.resource}}ByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := common.ParseAndValidateParamID(c.Param("id"), h.Val)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	var args {{.group}}.Update{{.resource}}sArgs
	if err := c.ShouldBindQuery(args); err != nil {
		common.ResponseError(c, err)
		return
	}

	{{.detail_prefix}}Details, err := h.Service.UpdateByID(ctx, id, args)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, {{.detail_prefix}}Details)
}

func (h RouterHandle) Delete{{.resource}}ByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := common.ParseAndValidateParamID(c.Param("id"), h.Val)
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	{{.detail_prefix}}Details, err := h.Service.Delete(ctx, id)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, {{.detail_prefix}}Details)
}

{{- end}}
