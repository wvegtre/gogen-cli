package database

import (
	"context"

	"gen-templates/internal/app/database/tables"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var InsertRow = func(ctx context.Context, gdb *gorm.DB, insertRow tables.ParentTable) (int64, error) {
	result := gdb.Table(insertRow.TableName()).Create(insertRow)
	if err := result.Error; err != nil {
		return 0, errors.WithStack(err)
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("not row affected after create")
	}
	return result.RowsAffected, nil
}

var UpdateRowByID = func(
	ctx context.Context, gdb *gorm.DB, id int64, updateRow tables.ParentTable,
) (int64, error) {
	result := gdb.Table(updateRow.TableName()).Where("id = ?", id).Updates(updateRow)
	if err := result.Error; err != nil {
		return 0, errors.WithStack(err)
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("not row affected after update")
	}
	return result.RowsAffected, nil
}

var DeleteRow = func(ctx context.Context, gdb *gorm.DB, id int64, deleteRow tables.ParentTable) (int64, error) {
	result := gdb.Table(deleteRow.TableName()).Delete(deleteRow, id)
	if err := result.Error; err != nil {
		return 0, errors.WithStack(err)
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("not row affected after delete")
	}
	return result.RowsAffected, nil
}

var QueryByID = func(ctx context.Context, gdb *gorm.DB, id int64, model tables.ParentTable) error {
	result := gdb.Table(model.TableName()).First(model, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.WithStack(err)
	}
	return nil
}

var CustomerQuery = func(ctx context.Context, gdb *gorm.DB, model tables.ParentTable) error {
	result := gdb.Find(model)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.WithStack(err)
	}
	return nil
}

var QueryList = func(
	ctx context.Context, gdb *gorm.DB,
	tableName string, model interface{},
	whereArgs map[string]interface{},
	options ...QueryOption,
) error {
	opArgs := &queryOptionArgs{}
	for _, op := range options {
		op(opArgs)
	}
	query := gdb.Table(tableName).Where(whereArgs)
	query = opArgs.setArgsToQuery(query)
	result := query.Find(model)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.WithStack(err)
	}
	return nil
}
