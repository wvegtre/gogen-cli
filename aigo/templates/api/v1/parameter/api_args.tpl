package parameter

import (
	"gen-templates/api/common"
	"gen-templates/internal/app/services"
)

type List{{.TableName}}Args struct {
    {{range $i, $v := .ListArgsRow}}{{$v}}
    {{end}}
	common.DefaultPaginationArgs
}

func (a List{{.TableName}}Args) SetDefaultPagination() List{{.TableName}}Args {
	if a.Page == 0 {
		a.Page = 1
	}
	if a.Size == 0 {
		a.Size = 100
	}
	return a
}

func (a List{{.TableName}}Args) ConvertToServiceArgs() services.List{{.TableName}}Args {
	return services.List{{.TableName}}Args{
	   {{range $i, $v := .ListArgsConvert}} {{$v}}: a.{{$v}},
       {{end}}
	}
}

