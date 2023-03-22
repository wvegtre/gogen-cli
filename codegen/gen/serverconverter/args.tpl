{{define "convert_args"}}
type List{{.service_prefix}}Args struct {
{{range $i, $v := .args_slice}}{{$v}}
{{end}}
}

func (a List{{.service_prefix}}Args) toDbQueryArgs() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	{{range $k, $v := .args_map}}m["{{$k}}"] = a.{{$v}}
	{{end}}
	return m
}

type Create{{.service_prefix}}Args struct {
{{range $i, $v := .args_slice}}{{if eq $v "Id"}}{{else}} {{$v}} {{end}}
{{end}}
}

type Update{{.service_prefix}}Args struct {
{{range $i, $v := .args_slice}}{{$v}}
{{end}}
}

func (a Update{{.service_prefix}}Args) toDbUpdateArgs() {{.group}}.{{.model}} {
	return {{.group}}.{{.model}}{
	{{range $k, $v := .args_map}}{{$v}}: a.{{$v}},
    {{end}}
	}
}

{{- end}}
