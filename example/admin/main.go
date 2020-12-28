package main

import (
	"github.com/cbwfree/micro-game/app"
	"github.com/cbwfree/micro-game/example/libs/def"
	mgo "github.com/cbwfree/micro-game/store/mongo"
	rds "github.com/cbwfree/micro-game/store/redis"
	"github.com/cbwfree/micro-game/utils/log"
	"github.com/micro/go-micro/v2"
)

func main() {
	app.New(def.SrvAdminName, def.Version, mgo.Flags, rds.Flags)

	// 初始化
	app.Init(
		micro.BeforeStart(beforeStart),
		micro.AfterStart(afterStart),
		micro.BeforeStop(beforeStop),
		micro.AfterStop(afterStop),
	)

	// 注册RPC服务
	app.AddHandler()

	// 启动服务
	if err := app.Run(); err != nil {
		log.Fatal("%+v", err)
	}
}

// beforeStart 启动之前执行
func beforeStart() error {
	// 连接Redis
	if err := rds.Connect(); err != nil {
		return err
	}
	// 连接MongoDB
	if err := mgo.Connect(); err != nil {
		return err
	}

	return nil
}

func afterStart() error {
	return nil
}

func beforeStop() error {
	return nil
}

func afterStop() error {
	// 关闭Redis
	if err := rds.Disconnect(); err != nil {
		log.Error("关闭Redis连接出错: %s", err.Error())
	}

	// 关闭MongoDB连接
	if err := mgo.Disconnect(); err != nil {
		log.Error("关闭MongoDB连接出错: %s", err.Error())
	}

	return nil
}
