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

func WithAllInOneFile(allInOneFile bool) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.AllInOneFile = allInOneFile
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

func WithSaveFileDefaultName(saveFileDefaultName string) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.SaveFileDefaultName = saveFileDefaultName
	}
}

func WithCharset(charset string) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.charset = charset
	}
}
