package logger

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

var loggerKey = ctxKey{}

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	once   sync.Once
	atom   zap.AtomicLevel
)

// InitLogger 는 전역 zap 로거를 초기화한다.
// logPath 가 비어 있으면 콘솔에만 기록한다.
// APP_ENV 환경 변수가 비어 있거나 "local"이 아닌 경우 지정된 logPath 파일과 콘솔에 로깅한다.
func InitLogger(logPath string) {
	once.Do(func() {
		env := os.Getenv("APP_ENV")
		var config zap.Config
		var options []zap.Option

		atom = zap.NewAtomicLevel()

		if env == "" || env == "local" {
			config = zap.NewDevelopmentConfig()
			config.Level = atom
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		} else {
			config = zap.NewProductionConfig()
			config.Level = atom
			config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

			outputs := []string{"stdout"}
			errOutputs := []string{"stderr"}

			if logPath != "" {
				outputs = append(outputs, logPath)
				errOutputs = append(errOutputs, logPath)
			}

			config.OutputPaths = outputs
			config.ErrorOutputPaths = errOutputs
		}

		var err error
		logger, err = config.Build(options...)
		if err != nil {
			panic(err)
		}
		sugar = logger.Sugar()
	})
}

// L 은 초기화된 전역 zap 로거를 반환한다.
func L() *zap.Logger {
	if logger == nil {
		InitLogger("")
	}
	return logger
}

// S 는 초기화된 전역 zap SugaredLogger를 반환한다.
func S() *zap.SugaredLogger {
	if sugar == nil {
		InitLogger("")
	}
	return sugar
}

// FromContext 는 컨텍스트에서 로거를 추출하거나 전역 로거를 반환한다.
func FromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return l
	}
	return L()
}

// WithContext 는 로거를 컨텍스트에 저장한다.
func WithContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

// SetLevel 은 로깅 레벨을 동적으로 변경한다.
func SetLevel(l zapcore.Level) {
	if logger == nil {
		InitLogger("")
	}
	atom.SetLevel(l)
}

// Named 는 로거 이름이 추가된 새로운 로거를 반환한다.
func Named(s string) *zap.Logger {
	return L().Named(s)
}

// With 는 필드가 추가된 새로운 로거를 반환한다.
func With(fields ...zap.Field) *zap.Logger {
	return L().With(fields...)
}

// Debug 는 디버그 레벨 로그를 기록한다.
func Debug(msg string, fields ...zap.Field) {
	L().Debug(msg, fields...)
}

// Info 는 정보 레벨 로그를 기록한다.
func Info(msg string, fields ...zap.Field) {
	L().Info(msg, fields...)
}

// Warn 는 경고 레벨 로그를 기록한다.
func Warn(msg string, fields ...zap.Field) {
	L().Warn(msg, fields...)
}

// Error 는 에러 레벨 로그를 기록한다.
func Error(msg string, fields ...zap.Field) {
	L().Error(msg, fields...)
}

// Fatal 는 치명적 에러 레벨 로그를 기록하고 프로그램을 종료한다.
func Fatal(msg string, fields ...zap.Field) {
	L().Fatal(msg, fields...)
}

// Sync 는 로거 버퍼를 비운다.
func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}
