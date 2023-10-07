package parameter

import (
	"gen-templates/api/common"
	"gen-templates/internal/app/services"
)

type List{{.TableName}}sArgs struct {
    {{range $i, $v := .ListArgsRow}}{{$v}}
    {{end}}
	common.DefaultPaginationArgs
}

func (a List{{.TableName}}sArgs) SetDefaultPagination() List{{.TableName}}sArgs {
	if a.Page == 0 {
		a.Page = 1
	}
	if a.Size == 0 {
		a.Size = 100
	}
	return a
}

func (a List{{.TableName}}sArgs) ConvertToServiceArgs() services.List{{.TableName}}sArgs {
	return services.List{{.TableName}}sArgs{
	   {{range $i, $v := .ListArgsConvert}} {{$v}}: a.{{$v}},
       {{end}}
	}
}

