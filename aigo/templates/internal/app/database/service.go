package database

type DBService service

type service interface {
	HookBeforeQuery() error
	GetTableName() string
}
