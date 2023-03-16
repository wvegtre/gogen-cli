package database

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// CustomerQuery A query function for special scenarios which query function can't cover, you should support a customer query
func (o Operation) CustomerQuery(
	ctx context.Context, query *gorm.DB, data interface{},
) error {
	return o.query(ctx, query, data)
}

// Query A query function that can cover most scenarios, it is recommended to use
func (o Operation) Query(
	ctx context.Context, s service,
	whereArgs map[string]interface{},
	model interface{}, options ...QueryOption,
) error {
	opArgs := &queryOptionArgs{}
	for _, op := range options {
		op(opArgs)
	}
	query := o.gdb.Table(s.GetTable())
	query = opArgs.setArgsToQuery(query)
	return o.query(ctx, query.Where(whereArgs), model)
}

func (o Operation) query(ctx context.Context, q *gorm.DB, model interface{}) error {
	result := q.Find(model)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.WithStack(err)
	}
	return nil
}

func (o Operation) QueryByID(ctx context.Context, s service, id int64, model interface{}) error {
	result := o.gdb.Table(s.GetTable()).First(model, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.WithStack(err)
	}
	return nil
}
