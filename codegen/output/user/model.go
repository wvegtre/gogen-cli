package user

import "gorm.io/gorm"

type UsersModel struct {
	gorm.Model
	Id        int    `gorm:"id" json:"id"`
	Name      string `gorm:"name" json:"name"`             // 昵称
	UserNo    string `gorm:"user_no" json:"user_no"`       // 用户编码，对外提供
	CreatedAt string `gorm:"created_at" json:"created_at"` // 创建时间
	UpdatedAt string `gorm:"updated_at" json:"updated_at"` // 更新时间
	DeletedAt string `gorm:"deleted_at" json:"deleted_at"` // 删除时间
}

func (UsersModel) TableName() string {
	return "users"
}

type UserAuthModel struct {
	gorm.Model
	Id      int    `gorm:"id" json:"id"`
	TokenId string `gorm:"token_id" json:"token_id"`
}

func (UserAuthModel) TableName() string {
	return "user_auth"
}
