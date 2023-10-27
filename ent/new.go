package ent

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"frame/conf"
)

func NewEntClient(lc fx.Lifecycle, log *zap.Logger, conf *conf.Config) *Client {
	client, err := Open("mysql", conf.DBDsn)
	if err != nil {
		log.Error("failed opening connection to mysql", zap.Error(err))
		panic(err)
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			client.Close()
			log.Info("ent client is closed")
			return nil
		},
	})
	return client

}
