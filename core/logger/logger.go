package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	MaxSize    = 10
	MaxBackups = 5
	Compress   = false
	channel    *Channel
)

// NewChannel Log channel
func NewChannel(c *Channel) *zap.Logger {
	channel = c

	// Get log writer
	writeSyncer := getLogWriter()

	// Get encoder
	encoder := getEncoder()

	// Log level
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(channel.Level)); err != nil {
		fmt.Println("init log level error")
		return nil
	}

	// New core
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// New logger
	logger := zap.New(core, zap.AddCaller())

	return logger
}

// Get encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// log format: NewJSONEncoder or NewConsoleEncoder
	if channel.Format == "console" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// Custom time encoder
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// Get log writer
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   channel.Path,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     channel.Days,
		Compress:   Compress,
	}

	// Additionally print to console.
	if channel.Console {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}

	return zapcore.AddSync(lumberJackLogger)
}
