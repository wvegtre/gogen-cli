package user

type UsersService struct{}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (u *UsersService) HookBeforeQuery() error {
	// TODO do something before run db query sql
	return nil
}

func (u UsersService) GetTable() string {
	return UsersModel{}.TableName()
}

type UserAuthService struct{}

func NewUserAuthService() *UserAuthService {
	return &UserAuthService{}
}

func (u *UserAuthService) HookBeforeQuery() error {
	// TODO do something before run db query sql
	return nil
}

func (u UserAuthService) GetTable() string {
	return UserAuthModel{}.TableName()
}
