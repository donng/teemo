package logging

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"strings"
	"time"
)

var Logger *zap.SugaredLogger

func init() {
	InitZapLogger()
}

func InitZapLogger() {
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.InfoLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	infoWriter := getWriter("runtime/logs/info.log")
	errorWriter := getWriter("runtime/logs/error.log")

	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(getEncoder(), zapcore.AddSync(errorWriter), errorLevel),
	)

	temp := zap.New(core, zap.AddCaller())

	Logger = temp.Sugar()

	defer Logger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "_%y%m%d.log", 1),
	)
	if err != nil {
		log.Fatalf("error getWriter err: %s", err)
	}
	return hook
}
