package user

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/database"
	"github.com/wvegtre/gogen-cli/output/easymoney/internal/app/database/user"
)

type UsersService struct {
	DBOperation    *database.Operation
	DBUsersService *user.UsersService
}

func NewUsersService() *UsersService {
	return &UsersService{
		DBOperation:    database.NewOperation(),
		DBUsersService: user.NewUsersService(),
	}
}

func (s *UsersService) Get(ctx context.Context, id int64) (user.UsersModel, error) {
	var dbUsersModel user.UsersModel
	err := s.DBOperation.QueryByID(ctx, s.DBUsersService, id, &dbUsersModel)
	if err != nil {
		return dbUsersModel, errors.WithStack(err)
	}
	return dbUsersModel, nil
}

func (s *UsersService) List(ctx context.Context, args ListUsersArgs, options ...database.QueryOption) ([]user.UsersModel, error) {
	var dbUsersModels []user.UsersModel
	whereArgs := args.toDbQueryArgs()
	err := s.DBOperation.Query(ctx, s.DBUsersService, whereArgs, &dbUsersModels, options...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dbUsersModels, nil
}

func (s *UsersService) Create(ctx context.Context, args CreateUsersArgs) (user.UsersModel, error) {
	var dbUsersModel user.UsersModel
	_, err := s.DBOperation.Create(ctx, s.DBUsersService, &dbUsersModel)
	if err != nil {
		return dbUsersModel, errors.WithStack(err)
	}
	return dbUsersModel, nil
}

func (s *UsersService) UpdateByID(ctx context.Context, id int64, args UpdateUsersArgs) (user.UsersModel, error) {
	var dbUsersModel user.UsersModel
	_, err := s.DBOperation.UpdateByID(ctx, s.DBUsersService, id, args.toDbUpdateArgs(), &dbUsersModel)
	if err != nil {
		return dbUsersModel, errors.WithStack(err)
	}
	return dbUsersModel, nil
}

func (s *UsersService) Delete(ctx context.Context, id int64) (user.UsersModel, error) {
	var dbUsersModel user.UsersModel
	_, err := s.DBOperation.DeleteByID(ctx, s.DBUsersService, id, &dbUsersModel)
	if err != nil {
		return dbUsersModel, errors.WithStack(err)
	}
	return dbUsersModel, nil
}

type UserAuthService struct {
	DBOperation       *database.Operation
	DBUserAuthService *user.UserAuthService
}

func NewUserAuthService() *UserAuthService {
	return &UserAuthService{
		DBOperation:       database.NewOperation(),
		DBUserAuthService: user.NewUserAuthService(),
	}
}

func (s *UserAuthService) Get(ctx context.Context, id int64) (user.UserAuthModel, error) {
	var dbUserAuthModel user.UserAuthModel
	err := s.DBOperation.QueryByID(ctx, s.DBUserAuthService, id, &dbUserAuthModel)
	if err != nil {
		return dbUserAuthModel, errors.WithStack(err)
	}
	return dbUserAuthModel, nil
}

func (s *UserAuthService) List(ctx context.Context, args ListUserAuthArgs) ([]user.UserAuthModel, error) {
	var dbUserAuthModels []user.UserAuthModel
	whereArgs := args.toDbQueryArgs()
	err := s.DBOperation.Query(ctx, s.DBUserAuthService, whereArgs, &dbUserAuthModels)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dbUserAuthModels, nil
}

func (s *UserAuthService) Create(ctx context.Context, args CreateUserAuthArgs) (user.UserAuthModel, error) {
	var dbUserAuthModel user.UserAuthModel
	_, err := s.DBOperation.Create(ctx, s.DBUserAuthService, &dbUserAuthModel)
	if err != nil {
		return dbUserAuthModel, errors.WithStack(err)
	}
	return dbUserAuthModel, nil
}

func (s *UserAuthService) UpdateByID(ctx context.Context, id int64, args UpdateUserAuthArgs) (user.UserAuthModel, error) {
	var dbUserAuthModel user.UserAuthModel
	_, err := s.DBOperation.UpdateByID(ctx, s.DBUserAuthService, id, args.toDbUpdateArgs(), &dbUserAuthModel)
	if err != nil {
		return dbUserAuthModel, errors.WithStack(err)
	}
	return dbUserAuthModel, nil
}

func (s *UserAuthService) Delete(ctx context.Context, id int64) (user.UserAuthModel, error) {
	var dbUserAuthModel user.UserAuthModel
	_, err := s.DBOperation.DeleteByID(ctx, s.DBUserAuthService, id, &dbUserAuthModel)
	if err != nil {
		return dbUserAuthModel, errors.WithStack(err)
	}
	return dbUserAuthModel, nil
}
