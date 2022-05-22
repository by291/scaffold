package init

// the entry of package init

// Init the application
func Init() {
	// init viper first to read all config
	err := ViperInit()
	if err != nil {
		panic(err)
	}

	// then init log, so we can use the config to customize zap logger
	err = LoggerInit()
	if err != nil {
		panic(err)
	}

	// init db
	err = DBInit()
	if err != nil {
		panic(err)
	}

	// init gin
	GinInit()
}
