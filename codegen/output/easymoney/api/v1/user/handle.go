package user

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/wvegtre/gogen-cli/output/easymoney/api/common"
	"github.com/wvegtre/gogen-cli/output/easymoney/api/common/consts"
	"github.com/wvegtre/gogen-cli/output/easymoney/api/v1/user/model"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/server/user"

	"github.com/gin-gonic/gin"
)

type UserHandle struct {
	Service *user.UsersService
	Val     *validator.Validate
}

func (h UserHandle) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if err := h.Val.Var(id, consts.ValidatorTagRequired); err != nil {
		common.ResponseError(c, err)
		return
	}
	idInt, _ := strconv.ParseInt(id, 10, 0)
	userDetail, err := h.Service.Get(ctx, idInt)
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, userDetail)
}

func (h UserHandle) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	var args model.GetUsersArgs
	if err := c.ShouldBindQuery(args); err != nil {
		common.ResponseError(c, err)
		return
	}
	userDetails, err := h.Service.List(ctx, args.ConvertToServiceArgs())
	if err != nil {
		common.ResponseError(c, err)
		return
	}
	common.ResponseOK(c, model.UsersListResponse{
		Users:      userDetails,
		Pagination: common.Pagination{},
	})
}
