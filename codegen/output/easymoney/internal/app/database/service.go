package database

type DBService service

type service interface {
	// Gorm not hook for before query, but we need it.
	// For example, when you query a big table, check limit and offset parameter is necessary before query.
	HookBeforeQuery() error
	GetTable() string
}
