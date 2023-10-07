package services

import (
	"context"
    "time"

	"gen-templates/api/common/consts"
	"gen-templates/internal/app/database"
	"gen-templates/internal/app/database/tables"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type {{.TableName}}Service struct {
	gdb *gorm.DB
	validator *validator.Validate
}

func New{{.TableName}}Service(gdb *gorm.DB) {{.TableName}}Service {
	return {{.TableName}}Service{
		gdb: gdb,
		validator: validator.New(),
	}
}

func (s {{.TableName}}Service) GetByID(ctx context.Context, id int64) (*{{.TableName}}Model, error) {
	if err := s.validator.Var("id", consts.ValidatorTagRequired); err != nil {
		return nil, errors.WithStack(err)
	}
	var db{{.TableName}} tables.{{.TableName}}Model
	err := database.QueryByID(ctx, s.gdb, id, &db{{.TableName}})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return convertToServiceModel(db{{.TableName}}), nil
}

func (s {{.TableName}}Service) List(ctx context.Context, args List{{.TableName}}Args, page, limit int) ({{.TableName}}Models, error) {
	if err := s.validator.Struct(struct {
		List{{.TableName}}Args
		Page  int `validate:"required"`
		Limit int `validate:"required"`
	}{
		List{{.TableName}}Args: args,
		Page:         page,
		Limit:        limit,
	}); err != nil {
		return nil, errors.WithStack(err)
	}
	var db{{.TableName}}s []tables.{{.TableName}}Model
	whereArgs := args.toDBQueryWhereArgs()
	err := database.QueryList(
		ctx, s.gdb, tables.{{.TableName}}Model{}.TableName(),
		&db{{.TableName}}s, whereArgs,
		database.WithQueryPage(page),
		database.WithQueryLimit(limit),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return convertToServiceModels(db{{.TableName}}s), nil
}

func (s {{.TableName}}Service) Create(ctx context.Context, args Create{{.TableName}}Args) (*{{.TableName}}Model, error) {
	if err := s.validator.Struct(args); err != nil {
		return nil, errors.WithStack(err)
	}
	createRow := args.toDBCreateRow()
	_, err := database.InsertRow(ctx, s.gdb, &createRow)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return convertToServiceModel(createRow), nil
}

func (s {{.TableName}}Service) UpdateByID(ctx context.Context, id int64, args Update{{.TableName}}Args) (*{{.TableName}}Model, error) {
	if err := s.validator.Struct(struct {
		Update{{.TableName}}Args
		ID int64 `validate:"required"`
	}{
		Update{{.TableName}}Args: args,
		ID:             id,
	}); err != nil {
		return nil, errors.WithStack(err)
	}
	updateRow := args.toDBUpdateRow()
	_, err := database.UpdateRowByID(ctx, s.gdb, id, &updateRow)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return convertToServiceModel(updateRow), nil
}

func (s {{.TableName}}Service) DeleteByID(ctx context.Context, id int64) (*{{.TableName}}Model, error) {
	if err := s.validator.Var(id, consts.ValidatorTagRequired); err != nil {
		return nil, errors.WithStack(err)
	}
	var db{{.TableName}} tables.{{.TableName}}Model
	_, err := database.DeleteRow(ctx, s.gdb, id, &db{{.TableName}})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return convertToServiceModel(db{{.TableName}}), nil
}

type {{.TableName}}Models []*{{.TableName}}Model
type {{.TableName}}Model struct {
    {{range $i, $v := .ModelFields}}{{$v}}
    {{end}}}

func convertToServiceModels(dbModels []tables.{{.TableName}}Model) {{.TableName}}Models {
	var serviceModels {{.TableName}}Models
	for _, v := range dbModels {
		serviceModels = append(serviceModels, convertToServiceModel(v))
	}
	return serviceModels
}

func convertToServiceModel(dbModel tables.{{.TableName}}Model) *{{.TableName}}Model {
	return &{{.TableName}}Model{
		{{range $i, $v := .ModelConvert}} {{$v}}: dbModel.{{$v}},
        {{end}}}
}

type List{{.TableName}}sArgs struct {
    {{range $i, $v := .ListArgsRow}}{{$v}}
    {{end}}}

func (a List{{.TableName}}Args) toDBQueryWhereArgs() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	{{range $key, $value := .ListArgMap}}m["{{$key}}"] = a.{{$value}}
    {{end}}return m
}

type Create{{.TableName}}Args struct {
    {{range $i, $v := .CreateArgsRow}}{{$v}}
    {{end}}}

func (a Create{{.TableName}}Args) toDBCreateRow() tables.{{.TableName}}Model {
	return tables.{{.TableName}}Model{
		{{range $i, $v := .CreateArgsConvert}} {{$v}}: a.{{$v}},
        {{end}}CreatedAt: 	time.Now().UTC(),
        UpdatedAt: 	time.Now().UTC(),
	}
}

type Update{{.TableName}}Args struct {
    {{range $i, $v := .UpdateArgsRow}}{{$v}}
    {{end}}}

func (a Update{{.TableName}}Args) toDBUpdateRow() tables.{{.TableName}}Model {
	return tables.{{.TableName}}Model{
	    {{range $i, $v := .UpdateArgsConvert}} {{$v}}: a.{{$v}},
        {{end}}UpdatedAt: 	time.Now().UTC(),
	}
}