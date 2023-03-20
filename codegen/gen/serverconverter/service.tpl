package user

import (
	"context"

	"echo-shopping/internal/app/database"
	"echo-shopping/internal/app/database/user"

	"github.com/pkg/errors"
)

type UserService service

type service struct {
	DBOperation   *database.Operation
	DBUserService *user.UserService
}

func NewService() *UserService {
	return &UserService{
		DBOperation:   database.NewOperation(),
		DBUserService: user.NewUserService(),
	}
}

func (s *UserService) Get(ctx context.Context, id int64) (user.UserModel, error) {
	var dbUser user.UserModel
	err := s.DBOperation.QueryByID(ctx, s.DBUserService, id, &dbUser)
	if err != nil {
		return dbUser, errors.WithStack(err)
	}
	return dbUser, nil
}

func (s *UserService) List(ctx context.Context, args ListArgs) ([]user.UserModel, error) {
	var dbUsers []user.UserModel
	// TODO
	whereArgs := args.toDbQueryArgs()
	// TODO query options
	err := s.DBOperation.Query(ctx, s.DBUserService, whereArgs, &dbUsers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dbUsers, nil
}

func (s *UserService) Create(ctx context.Context, args CreateArgs) (user.UserModel, error) {
	var dbUser user.UserModel
	_, err := s.DBOperation.Create(ctx, s.DBUserService, &dbUser)
	if err != nil {
		return dbUser, errors.WithStack(err)
	}
	return dbUser, nil
}

func (s *UserService) UpdateByID(ctx context.Context, id int64, args UpdateArgs) (user.UserModel, error) {
	var dbUser user.UserModel
	_, err := s.DBOperation.UpdateByID(ctx, s.DBUserService, id, args.toDbUpdateArgs(), &dbUser)
	if err != nil {
		return dbUser, errors.WithStack(err)
	}
	return dbUser, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) (user.UserModel, error) {
	var dbUser user.UserModel
	_, err := s.DBOperation.DeleteByID(ctx, s.DBUserService, id, &dbUser)
	if err != nil {
		return dbUser, errors.WithStack(err)
	}
	return dbUser, nil
}
