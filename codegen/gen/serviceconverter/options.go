package serviceconverter

type ServiceGenConfigOption func(config *serviceConverterConfig)

func WithSaveDir(saveDir string) ServiceGenConfigOption {
	return func(config *serviceConverterConfig) {
		config.SaveDir = saveDir
	}
}

func WithSavePackageName(packageName string) ServiceGenConfigOption {
	return func(config *serviceConverterConfig) {
		config.SavePackageName = packageName
	}
}

func WithSaveProjectName(saveProjectName string) ServiceGenConfigOption {
	return func(config *serviceConverterConfig) {
		config.SaveProjectName = saveProjectName
	}
}
