package init

import (
	"by291.fun/scaffold/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func LoggerInit() error {
	logConf := &global.Config.Log

	// unmarshal log level
	level := new(zapcore.Level)
	err := level.UnmarshalText([]byte(logConf.Level))
	if err != nil {
		return errors.Wrap(err, "unmarshal log level fail")
	}

	var core zapcore.Core
	if logConf.Mode == "dev" {
		encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
	} else if logConf.Mode == "prod" {
		// lumberjack syncer
		syncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConf.Filename,
			MaxSize:    logConf.MaxSize, // MB
			MaxBackups: logConf.MaxBackups,
			MaxAge:     logConf.MaxAge, // days
			LocalTime:  true,
		})

		encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		core = zapcore.NewCore(encoder, syncer, level)
	}

	logger := zap.New(core)
	zap.ReplaceGlobals(logger)

	return nil
}
