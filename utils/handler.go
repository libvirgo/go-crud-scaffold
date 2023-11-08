package utils

import (
	"frame/conf"
	"frame/ent"
	"frame/utils/redislock"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Void struct {
}

type Router interface {
}

func AsRouter(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Router)),
		fx.ResultTags(`group:"routers"`),
	)
}

func NewHandler(logger *zap.Logger,
	e *ent.Client,
	locker redislock.Locker,
	rd *redis.Client,
	conf *conf.Config,
	queue *asynq.Client,
) Handler {
	var entClient *ent.Client
	if conf.Debug {
		entClient = e.Debug()
	} else {
		entClient = e
	}
	return Handler{
		Ent:    entClient,
		Redis:  rd,
		Locker: locker,
		Config: conf,
		Logger: logger,
		Queue:  queue,
	}
}

type Handler struct {
	Ent    *ent.Client
	Redis  *redis.Client
	Locker redislock.Locker
	Config *conf.Config
	Queue  *asynq.Client
	*zap.Logger
}

type ResponseVO struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type StatusVO struct {
	Status string `json:"status"`
}
