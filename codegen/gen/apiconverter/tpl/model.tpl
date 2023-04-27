{{define "convert_handle"}}
type Get{{.resource}}sArgs struct {
	Name  string `form:"name,omitempty"`
	Email string `form:"email,omitempty"`
	common.DefaultPaginationArgs
}

func (a *Get{{.resource}}sArgs) SetDefaultPagination() {
	if a.Page == 0 {
		a.Page = 1
	}
	if a.Limit == 0 {
		a.Limit = 100
	}
}

func (a *Get{{.resource}}sArgs) ConvertToServiceArgs() {{.group}}.List{{.resource}}sArgs {
	return {{.group}}.List{{.resource}}sArgs{}
}

func (a *Get{{.resource}}sArgs) GenQueryOptions() []database.QueryOption {
	if a.Page == 0 || a.Limit == 0 {
		return nil
	}
	return []database.QueryOption{
		database.WithQueryLimit(a.Limit),
		database.WithQueryPage(a.Page),
	}
}

type {{.resource}}sListResponse struct {
	{{.resource}}s      []{{.db_prefix}}Db.{{.model}} `json:"users"`
	Pagination common.Pagination   `json:"pagination"`
}

{{- end}}
