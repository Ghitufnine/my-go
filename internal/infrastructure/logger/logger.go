package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() (*zap.Logger, error) {

	encoderCfg := zap.NewProductionEncoderConfig()

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.LevelKey = "level"
	encoderCfg.MessageKey = "message"
	encoderCfg.CallerKey = "caller"

	encoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())

	return logger, nil
}
