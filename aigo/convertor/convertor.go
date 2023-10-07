package convertor

import (
	"context"

	"gogen-cli/aigo/model"
)

type Processor interface {
	GetTableFields(ctx context.Context, tableName string) (model.TableFields, error)
}
