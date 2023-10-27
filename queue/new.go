package queue

import (
	"context"
	"frame/conf"
	"frame/utils"
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type H utils.Handler

func New() fx.Option {
	return fx.Module("queue",
		fx.Provide(NewClient),
		fx.Provide(NewServer),
		fx.Provide(NewScheduler),
		fx.Invoke(func(server *asynq.Server) {}),
		fx.Invoke(func(scheduler *asynq.Scheduler) {}),
	)
}

func NewClient(lc fx.Lifecycle, conf *conf.Config) *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: conf.Redis.Addr, Password: conf.Redis.Password})
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return client.Close()
		},
	})
	return client
}

func NewServer(lc fx.Lifecycle, logger *zap.Logger, conf *conf.Config, h utils.Handler) *asynq.Server {
	srv := asynq.NewServer(asynq.RedisClientOpt{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
	}, asynq.Config{
		Concurrency: 10,
		Logger:      logger.Sugar(),
	})
	_ = H(h)
	mux := asynq.NewServeMux()
	{
	}
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				logger.Info("asynq run")
				err := srv.Run(mux)
				if err != nil {
					logger.Error("asynq run error", zap.String("error", err.Error()))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown()
			return nil
		},
	})
	return srv
}

func NewScheduler(lc fx.Lifecycle, logger *zap.Logger, conf *conf.Config) *asynq.Scheduler {
	srv := asynq.NewScheduler(asynq.RedisClientOpt{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
	}, &asynq.SchedulerOpts{
		Logger: logger.Sugar(),
		// Location specifies the time zone to use when scheduling tasks.
		Location: utils.PalestineZone,
	})
	//var entryId string
	//var err error
	//srv.Register("@every 1h", asynq.NewTask())
	//if err != nil {
	//	logger.Error("asynq scheduler register error", zap.String("error", err.Error()))
	//}
	//logger.Info("asynq scheduler register", zap.String("entryId", entryId))
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				logger.Info("asynq scheduler run")
				err := srv.Run()
				if err != nil {
					logger.Error("asynq scheduler run error", zap.String("error", err.Error()))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown()
			return nil
		},
	})
	return srv
}
