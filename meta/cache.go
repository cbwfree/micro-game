package meta

import (
	"context"
	rds "github.com/cbwfree/micro-game/store/redis"
	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
)

func Client() *redis.Client {
	if client != nil {
		return client
	}
	return rds.Client()
}

func SetCacheClient(c *redis.Client) {
	client = c
}

// LoadMetaCache 获取Meta缓存
func LoadMetaCache(key string) (map[string]string, error) {
	res, err := Client().HGetAll(context.TODO(), key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, nil
}

// MultiLoadMetaCache 批量获取Meta缓存
func MultiLoadMetaCache(keys ...string) (map[string]map[string]string, error) {
	ctx := context.TODO()
	cmds, err := Client().Pipelined(ctx, func(pip redis.Pipeliner) error {
		for _, key := range keys {
			pip.HGetAll(ctx, key)
		}
		return nil
	})
	if err != nil && err != redis.Nil {
		return nil, err
	}

	var result = make(map[string]map[string]string)
	for i, key := range keys {
		cmd, ok := cmds[i].(*redis.StringStringMapCmd)
		if !ok {
			continue
		}

		res, err := cmd.Result()
		if err != nil {
			return nil, err
		}

		result[key] = res
	}

	return result, nil
}

// SaveMetaCache 保存Meta缓存
func SaveMetaCache(key string, mt *Meta) error {
	return Client().HMSet(context.TODO(), key, rds.Args{}.AddFlat(mt.Metadata())...).Err()
}

// MultiSaveMetaCache 批量保存Meta缓存
func MultiSaveMetaCache(metas map[string]*Meta) error {
	ctx := context.TODO()
	_, err := Client().Pipelined(ctx, func(pip redis.Pipeliner) error {
		for key, mt := range metas {
			pip.HMSet(ctx, key, rds.Args{}.AddFlat(mt.Metadata())...)
		}
		return nil
	})
	return err
}

// DeleteMeta 删除meta信息
func DeleteMetaCache(key ...string) error {
	return Client().Del(context.TODO(), key...).Err()
}
