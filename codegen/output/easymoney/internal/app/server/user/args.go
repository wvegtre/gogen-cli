package user

import (
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/database/user"
)

type ListUsersArgs struct {
	Id        int
	Name      string
	UserNo    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (a ListUsersArgs) toDbQueryArgs() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	m["created_at"] = a.CreatedAt
	m["deleted_at"] = a.DeletedAt
	m["id"] = a.Id
	m["name"] = a.Name
	m["updated_at"] = a.UpdatedAt
	m["user_no"] = a.UserNo
	return m
}

type CreateUsersArgs struct {
	Id        int `json:"id,omitempty" form:"id,omitempty"`
	Name      string
	UserNo    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type UpdateUsersArgs struct {
	Id        int
	Name      string
	UserNo    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (a UpdateUsersArgs) toDbUpdateArgs() user.UsersModel {
	return user.UsersModel{
		CreatedAt: a.CreatedAt,
		DeletedAt: a.DeletedAt,
		Id:        a.Id,
		Name:      a.Name,
		UpdatedAt: a.UpdatedAt,
		UserNo:    a.UserNo,
	}
}

type ListUserAuthArgs struct {
	Id      int
	TokenId string
}

func (a ListUserAuthArgs) toDbQueryArgs() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	m["id"] = a.Id
	m["token_id"] = a.TokenId

	return m
}

type CreateUserAuthArgs struct {
	Id      int
	TokenId string
}

type UpdateUserAuthArgs struct {
	Id      int
	TokenId string
}

func (a UpdateUserAuthArgs) toDbUpdateArgs() user.UserAuthModel {
	return user.UserAuthModel{
		Id:      a.Id,
		TokenId: a.TokenId,
	}
}
