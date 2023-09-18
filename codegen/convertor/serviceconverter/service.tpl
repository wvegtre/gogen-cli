{{define "convert"}}
type {{.service_prefix}}Service struct{}

func New{{.service_prefix}}Service() *{{.service_prefix}}Service {
	return &{{.service_prefix}}Service{}
}

func (u *{{.service_prefix}}Service) HookBeforeQuery() error {
	// TODO do something before run db query sql
	return nil
}

func (u {{.service_prefix}}Service) GetTable() string {
	return {{.model}}{}.TableName()
}
{{- end}}