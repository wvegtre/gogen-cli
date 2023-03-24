package database

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Operation struct {
	gdb *gorm.DB
}

func NewOperation() *Operation {
	return &Operation{}
}

func (o Operation) Create(ctx context.Context, s service, data interface{}) (int64, error) {
	result := o.gdb.Table(s.GetTable()).Create(data)
	if err := result.Error; err != nil {
		return 0, errors.WithStack(err)
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("not row affected after create")
	}
	return result.RowsAffected, nil
}

func (o Operation) UpdateByID(ctx context.Context, s service, id int64, args interface{}, data interface{}) (int64, error) {
	db := o.gdb.Table(s.GetTable())
	// 默认是会用 s.GetDBModel() 返回的 struct 里边的 ID 字段，但是比较隐性，为了防止误用，这里明确指定要更新的 ID
	result := db.Where("id = ?", id).Updates(data)
	if err := result.Error; err != nil {
		return 0, errors.WithStack(err)
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("not row affected after update")
	}
	return result.RowsAffected, nil
}

func (o Operation) DeleteByID(ctx context.Context, s service, id int64, data interface{}) (int64, error) {
	result := o.gdb.Table(s.GetTable()).Delete(data, id)
	if err := result.Error; err != nil {
		return 0, errors.WithStack(err)
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("not row affected after delete")
	}
	return result.RowsAffected, nil
}
