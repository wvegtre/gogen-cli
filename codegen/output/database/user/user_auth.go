package user

import (
	"gorm.io/gorm"
)

type UserAuthModel struct {
	gorm.Model
	Id      int    `gorm:"id" json:"id"`
	TokenId string `gorm:"token_id" json:"token_id"`
}

func (UserAuthModel) TableName() string {
	return "user_auth"
}
