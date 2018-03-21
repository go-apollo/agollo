package agollo

var logger LoggerInterface

func initLogger(ILogger LoggerInterface) {
	logger = ILogger
}

type LoggerInterface interface {
	Debugf(format string, params ...interface{})

	Infof(format string, params ...interface{})

	Warnf(format string, params ...interface{})

	Errorf(format string, params ...interface{})

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{})

	Error(v ...interface{})
}

// func initSeeLog(configPath string) LoggerInterface {
// 	logger, err := seelog.LoggerFromConfigAsFile(configPath)

// 	//if error is happen change to default config.
// 	if err != nil {
// 		logger, err = seelog.LoggerFromConfigAsBytes([]byte("<seelog />"))
// 	}

// 	logger.SetAdditionalStackDepth(1)
// 	seelog.ReplaceLogger(logger)
// 	defer seelog.Flush()

// 	return logger
// }
