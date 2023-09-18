package convertor

type GenDBCodeParameter struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
	IP       string `validate:"required"`
	Port     string `validate:"required"`
	Database string `validate:"required"`
	Charset  string // default utf8
}

type GenDBServiceCodeParameter struct {
	Operator     ConnectOperator
	TargetTables string
	SaveDir      string
}
