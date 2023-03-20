package dbconverter

type MySQLConfigOption func(config *mySQLConverterConfig)

func WithTables(tables []string) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.tables = tables
	}
}

func WithEnableJsonTag(enableJsonTag bool) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.EnableJsonTag = enableJsonTag
	}
}

func WithSaveDir(saveDir string) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.SaveDir = saveDir
	}
}

func WithSaveFilePrefix(saveFilePrefix string) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.SaveFilePrefix = saveFilePrefix
	}
}

func WithCharset(charset string) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.charset = charset
	}
}
