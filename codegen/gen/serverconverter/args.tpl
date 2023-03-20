{{define "convert_args"}}
type List{{.table}}Args struct {
{{range $i, $v := .args_slice}}
    {{$v}}
{{end}}
}

func (a List{{.table}}Args) toDbQueryArgs() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	{{range $k, $v := .args_map}}
        m[{{$k}}] = {{$v}}
    {{end}}
	return m
}

type Create{{.table}}Args struct {
{{range $i, $v := .args_slice}}
    {{if eq $v "id"}}{{else}} {{$v}} {{end}}
{{end}}
}

type Update{{.table}}Args struct {
{{range $i, $v := .args_slice}}
    {{$v}}
{{end}}
}

func (a Update{{.table}}Args) toDbUpdateArgs() {{.group}}.{{.table}}Model {
	// TODO service args to db args
	return user.UserModel{}
}
