package db

type ConnectOperator interface {
	GenStructByDBFields(parameter GenParameter) error
	GenServiceForDBStruct() error
}

func NewMySQLDBConnect() ConnectOperator {
	return &MySQLDBConnect{}
}
