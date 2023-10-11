package tables

type {{.TableName}}Model struct {
{{range $i, $v := .StructFields}}{{$v}}
{{end}}
}

func ({{.TableName}}Model) TableName() string {
	return "{{.DBTableName}}"
}

