package serviceconverter

type ServiceGenConfigOption func(config *serviceConverterConfig)

func WithSaveDir(saveDir string) ServiceGenConfigOption {
	return func(config *serviceConverterConfig) {
		config.SaveDir = saveDir
	}
}

//
//func WithSaveFilePrefix(saveFilePrefix string) ServiceGenConfigOption {
//	return func(config *serviceConverterConfig) {
//		config.SaveFilePrefix = saveFilePrefix
//	}
//}

func WithSaveFileDefaultName(saveFileDefaultName string) ServiceGenConfigOption {
	return func(config *serviceConverterConfig) {
		config.SaveFileDefaultName = saveFileDefaultName
	}
}
