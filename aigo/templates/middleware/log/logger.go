package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	bizLogKey = "biz_log_fields"
)

var (
	bizLog = new(Logger)
)

func init() {
	zapLogger, err := newZap()
	if err != nil {
		panic("create zap logger error" + err.Error())
	}
	bizLog.logger = zapLogger
}

func Get() *Logger {
	return bizLog
}

func GetBizLoggerKey() string {
	return bizLogKey
}

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) InfoCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, appendMoreFieldsFromContext(ctx)...)
	l.logger.Info(msg, fields...)
}

func (l *Logger) WarnCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, appendMoreFieldsFromContext(ctx)...)
	l.logger.Warn(msg, fields...)
}

func (l *Logger) ErrorCtx(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, appendMoreFieldsFromContext(ctx)...)
	l.logger.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.logger.DPanic(msg, fields...)
}

func newZap() (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "event_time"
	encoderConfig.LevelKey = "severity"
	encoderConfig.MessageKey = "message"
	encoderConfig.StacktraceKey = "stack_trace"
	encoderConfig.EncodeLevel = levelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.EncoderConfig = encoderConfig
	cfg.Encoding = "json"
	cfg.Sampling = defaultSamplingConfig()

	level, err := zapcore.ParseLevel("info")
	if err != nil {
		return nil, err
	}
	cfg.Level = zap.NewAtomicLevelAt(level)

	return cfg.Build(zap.AddStacktrace(zapcore.DPanicLevel), zap.AddCallerSkip(0))
}

// levelEncoder serializes a Level to a Google Operations Suite logging severity.
func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l { // nolint
	case zapcore.WarnLevel:
		enc.AppendString("WARNING")
	case zapcore.DPanicLevel, zapcore.PanicLevel:
		enc.AppendString("CRITICAL")
	case zapcore.FatalLevel:
		enc.AppendString("EMERGENCY")
	default:
		enc.AppendString(l.CapitalString())
	}
}

func defaultSamplingConfig() *zap.SamplingConfig {
	return &zap.SamplingConfig{
		Initial:    64, // log 64 entries per second before sampling
		Thereafter: 10, // log at 10th entry after exceeding the 64 entries
	}
}

func appendMoreFieldsFromContext(ctx context.Context) []zap.Field {
	if ctx == nil {
		return nil
	}

	ctxLoggerFields, ok := ctx.Value(bizLogKey).([]zap.Field)
	if !ok {
		return nil
	}

	return ctxLoggerFields
}
