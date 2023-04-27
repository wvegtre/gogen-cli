package user

import (
	"github.com/wvegtre/gogen-cli/output/easymoney/api/common"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/database"
	userDb "github.com/wvegtre/gogen-cli/output/easymoney/internal/app/database/user"
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

func (a *GetUsersArgs) ConvertToServiceArgs() user.ListUsersArgs {
	return user.ListUsersArgs{}
}

func (a *GetUsersArgs) GenQueryOptions() []database.QueryOption {
	if a.Page == 0 || a.Limit == 0 {
		return nil
	}
	return []database.QueryOption{
		database.WithQueryLimit(a.Limit),
		database.WithQueryPage(a.Page),
	}
}

type UsersListResponse struct {
	Users      []userDb.UsersModel `json:"users"`
	Pagination common.Pagination   `json:"pagination"`
}
