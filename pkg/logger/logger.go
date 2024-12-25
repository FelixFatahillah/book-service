package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Encoder
			zapcore.AddSync(os.Stdout),            // Console output
			zapcore.InfoLevel,                     // Log level
		),
	)

	//Log, err = config.Build(zap.AddCallerSkip(1))
	Log = zap.New(core, zap.AddCallerSkip(1), zap.AddCaller())

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	Log.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Log.Error(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	Log.Warn(message, fields...)
}
