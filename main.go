package main

import (
	"frame/conf"
	"frame/ent"
	"frame/queue"
	"frame/redislock"
	"frame/router"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(conf *conf.Config) *zap.Logger {
	if conf.Debug {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, _ := config.Build()
		return logger
	} else {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		logger, _ := config.Build()
		return logger
	}
}

func NewRedis(conf *conf.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       0,
	})
	return rdb
}

func main() {
	fx.New(
		fx.Provide(
			conf.NewConfig,
			NewLogger,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			ent.NewEntClient,
			NewRedis,
			redislock.NewRedisLocker,
		),
		router.New(),
		queue.New(),
	).Run()
}
