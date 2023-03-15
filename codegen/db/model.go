package db

type GenParameter struct {
	Operator    ConnectOperator
	UserName    string `validate:"required"`
	Password    string `validate:"required"`
	IP          string `validate:"required"`
	Port        string `validate:"required"`
	Database    string `validate:"required"`
	Charset     string // default utf8
	TargetTable string
}
