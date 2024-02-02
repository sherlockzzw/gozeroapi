package config

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Etcd struct {
		Hosts []string `json:"hosts"`
		Key   string   `json:"key"`
	}
	RedisX string
}

type RedisX struct {
	Addr     string `json:"addr"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Db       int    `json:"db"`
}

func getRedis(config Config) (redisX *RedisX) {
	redisResp, err := EtcdClient.Get(context.Background(), config.RedisX)
	if err != nil {
		panic(err)
	}
	redisX = &RedisX{}
	if err = json.Unmarshal(redisResp.Kvs[0].Value, redisX); err != nil {
		panic(err)
	}

	return
}
