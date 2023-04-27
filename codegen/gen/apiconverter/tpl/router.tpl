{{define "convert_router"}}
type {{.resource}}sRouter struct {
}

func (r {{.resource}}sRouter) Init() *router.Router {

	h := RouterHandle{
		Service: {{.group}}.New{{.service_prefix}}Service(),
	}
	parent := router.NewRouterWithPath("/{{.path_prefix}}")
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, "", h.Get{{.resource}}s))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodGet, ":id", h.Get{{.resource}}ByID))
	parent.AddSubRouterGroup(router.NewRouter(http.MethodPost, "", h.Create{{.resource}}))
    parent.AddSubRouterGroup(router.NewRouter(http.MethodPatch, ":id", h.Update{{.resource}}ByID))
	return parent
}

{{- end}}
