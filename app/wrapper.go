package app

import (
	"context"
	"github.com/cbwfree/micro-game/utils/errors"
	"github.com/cbwfree/micro-game/utils/log"
	"github.com/micro/go-micro/v2/server"
	"time"
)

func serverWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		now := time.Now()
		err := fn(ctx, req, rsp)

		if Opts.Dev {
			if err != nil {
				if errors.IsStack() {
					log.Error("[Server] received [%s], uptime: %s, error: %+v", req.Endpoint(), time.Since(now), err)
				} else {
					log.Error("[Server] received [%s], uptime: %s, error: %s", req.Endpoint(), time.Since(now), err)
				}
			} else {
				log.Info("[Server] received [%s], uptime: %s", req.Endpoint(), time.Since(now))
			}
		}

		if err == nil {
			return nil
		}

		return errors.MicroError(Id(), err)
	}
}

func subscriberWrapper(fn server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		now := time.Now()
		err := fn(ctx, msg)

		if Opts.Dev {
			if err != nil {
				if errors.IsStack() {
					log.Error("[Subscriber] received [%s], uptime: %s, error: %+v", msg.Topic(), time.Since(now), err)
				} else {
					log.Error("[Subscriber] received [%s], uptime: %s, error: %s", msg.Topic(), time.Since(now), err)
				}
			} else {
				log.Info("[Subscriber] received [%s], uptime: %s", msg.Topic(), time.Since(now))
			}
		}

		if err == nil {
			return nil
		}

		return errors.MicroError(Id(), err)
	}
}
