package model

import (
	"github.com/wvegtre/gogen-cli/output/easymoney/api/common"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/database/user"
)

type UsersListResponse struct {
	Users      []user.UsersModel `json:"users"`
	Pagination common.Pagination `json:"pagination"`
}
