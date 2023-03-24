package model

import (
	"github.com/wvegtre/gogen-cli/output/easymoney/api/common"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/server/user"
)

type GetUsersArgs struct {
	Name  string `form:"name,omitempty"`
	Email string `form:"email,omitempty"`
	common.DefaultPaginationArgs
}

func (a *GetUsersArgs) SetDefaultPagination() {
	if a.Page == 0 {
		a.Page = 1
	}
	if a.Limit == 0 {
		a.Limit = 100
	}
}

func (a GetUsersArgs) ConvertToServiceArgs() user.ListUsersArgs {
	return user.ListUsersArgs{}
}
