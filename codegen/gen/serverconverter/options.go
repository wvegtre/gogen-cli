package serverconverter

type ServerGenConfigOption func(config *serverConverterConfig)

func WithSaveDir(saveDir string) ServerGenConfigOption {
	return func(config *serverConverterConfig) {
		config.SaveDir = saveDir
	}
}

//
//func WithSaveFilePrefix(saveFilePrefix string) ServerGenConfigOption {
//	return func(config *serverConverterConfig) {
//		config.SaveFilePrefix = saveFilePrefix
//	}
//}

func WithSaveFileDefaultName(saveFileDefaultName string) ServerGenConfigOption {
	return func(config *serverConverterConfig) {
		config.SaveFileDefaultName = saveFileDefaultName
	}
}
