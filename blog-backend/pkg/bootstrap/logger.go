package bootstrap

import (
	"os"

	"github.com/sirupsen/logrus"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	aliyunLogger "github.com/go-kratos/kratos/contrib/log/aliyun/v2"
	fluentLogger "github.com/go-kratos/kratos/contrib/log/fluent/v2"
	logrusLogger "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	tencentLogger "github.com/go-kratos/kratos/contrib/log/tencent/v2"
	zapLogger "github.com/go-kratos/kratos/contrib/log/zap/v2"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"kratos-blog/gen/api/go/common/conf"
)

type LoggerType string

const (
	LoggerTypeStd     LoggerType = "std"
	LoggerTypeFluent  LoggerType = "fluent"
	LoggerTypeLogrus  LoggerType = "logrus"
	LoggerTypeZap     LoggerType = "zap"
	LoggerTypeAliyun  LoggerType = "aliyun"
	LoggerTypeTencent LoggerType = "tencent"
)

func NewLoggerProvider(loggerType LoggerType, cfg *conf.Logger, serviceInfo *ServiceInfo) log.Logger {
	l := createLogger(loggerType, cfg)

	return log.With(
		l,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

func createLogger(loggerType LoggerType, cfg *conf.Logger) log.Logger {
	switch loggerType {
	default:
		fallthrough
	case LoggerTypeStd:
		return createStdLogger()
	case LoggerTypeFluent:
		return createFluentLogger(cfg)
	case LoggerTypeZap:
		return createZapLogger(cfg)
	case LoggerTypeLogrus:
		return createLogrusLogger(cfg)
	case LoggerTypeAliyun:
		return createAliyunLogger(cfg)
	case LoggerTypeTencent:
		return createTencentLogger(cfg)
	}
}

func createStdLogger() log.Logger {
	l := log.NewStdLogger(os.Stdout)
	return l
}

type zapWriteSyncer struct {
	output []string
}

func (x *zapWriteSyncer) Write(p []byte) (n int, err error) {
	x.output = append(x.output, string(p))
	return len(p), nil
}

func (x *zapWriteSyncer) Sync() error {
	return nil
}

func createZapLogger(cfg *conf.Logger) log.Logger {
	syncer := &zapWriteSyncer{}
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     cfg.Zap.MessageKey,
		LevelKey:       cfg.Zap.LevelKey,
		NameKey:        cfg.Zap.NameKey,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), syncer, zap.DebugLevel)
	options := zap.New(core).WithOptions()
	logger := zapLogger.NewLogger(options)

	return logger
}

func createLogrusLogger(cfg *conf.Logger) log.Logger {
	loggerLevel, err := logrus.ParseLevel(cfg.Logrus.Level)
	if err != nil {
		loggerLevel = logrus.InfoLevel
	}

	var loggerFormatter logrus.Formatter
	switch cfg.Logrus.Formatter {
	default:
		fallthrough
	case "text":
		loggerFormatter = &logrus.TextFormatter{
			DisableColors:    cfg.Logrus.DisableColors,
			DisableTimestamp: cfg.Logrus.DisableTimestamp,
			TimestampFormat:  cfg.Logrus.TimestampFormat,
		}
		break
	case "json":
		loggerFormatter = &logrus.JSONFormatter{
			DisableTimestamp: cfg.Logrus.DisableTimestamp,
			TimestampFormat:  cfg.Logrus.TimestampFormat,
		}
		break
	}

	logger := logrus.New()
	logger.Level = loggerLevel
	logger.Formatter = loggerFormatter

	wrapped := logrusLogger.NewLogger(logger)
	return wrapped
}

func createFluentLogger(cfg *conf.Logger) log.Logger {
	logger, err := fluentLogger.NewLogger(cfg.Fluent.Endpoint)
	if err != nil {
		panic("create fluent logger failed")
		return nil
	}
	return logger
}

func createAliyunLogger(cfg *conf.Logger) log.Logger {
	logger := aliyunLogger.NewAliyunLog(
		aliyunLogger.WithProject(cfg.Aliyun.Project),
		aliyunLogger.WithEndpoint(cfg.Aliyun.Endpoint),
		aliyunLogger.WithAccessKey(cfg.Aliyun.AccessKey),
		aliyunLogger.WithAccessSecret(cfg.Aliyun.AccessSecret),
	)
	return logger
}

func createTencentLogger(cfg *conf.Logger) log.Logger {
	logger, err := tencentLogger.NewLogger(
		tencentLogger.WithTopicID(cfg.Tencent.TopicId),
		tencentLogger.WithEndpoint(cfg.Tencent.Endpoint),
		tencentLogger.WithAccessKey(cfg.Tencent.AccessKey),
		tencentLogger.WithAccessSecret(cfg.Tencent.AccessSecret),
	)
	if err != nil {
		panic(err)
		return nil
	}
	return logger
}
