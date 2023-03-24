package serverconverter

type ServerGenConfigOption func(config *serverConverterConfig)

func WithSaveDir(saveDir string) ServerGenConfigOption {
	return func(config *serverConverterConfig) {
		config.SaveDir = saveDir
	}
}

func WithSavePackageName(packageName string) ServerGenConfigOption {
	return func(config *serverConverterConfig) {
		config.SavePackageName = packageName
	}
}

func WithSaveProjectName(saveProjectName string) ServerGenConfigOption {
	return func(config *serverConverterConfig) {
		config.SaveProjectName = saveProjectName
	}
}
