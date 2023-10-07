package detector

import (
	"context"

	"gogen-cli/aigo/model"
)

type Handler interface {
	GetTableFields(ctx context.Context, tableName string) (model.TableFields, error)
}
