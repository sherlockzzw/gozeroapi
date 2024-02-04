package config

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"os"
	"path/filepath"
	"time"
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
	redisResp, err := etcdClient.Get(context.Background(), config.RedisX)
	if err != nil {
		panic(err)
	}
	redisX = &RedisX{}
	if err = json.Unmarshal(redisResp.Kvs[0].Value, redisX); err != nil {
		panic(err)
	}

	return
}

func Init() (config *Config) {
	config = new(Config)
	configOnce.Do(func() {
		var (
			filePath string
			err      error
		)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				panic("read config yaml file failed")
			case <-time.After(time.Nanosecond):
				env := os.Getenv("environ")
				switch env {
				case service.ProMode:
					filePath, err = filepath.Abs("etc/release.apiclipfilm.yaml")
					if err != nil {
						panic(err)
					}
				case service.TestMode:
					filePath, err = filepath.Abs("etc/test.apiclipfilm.yaml")
					if err != nil {
						panic(err)
					}
				default:
					filePath, err = filepath.Abs("etc/dev.apiclipfilm.yaml")
					if err != nil {
						panic(err)
					}
				}

				conf.MustLoad(filePath, config)

				logx.WithContext(ctx).Infof("successfully set config, path: %s, env: %s", filePath, env)
				RegisterEtcd(ctx, config)
				return
			}
		}
	})

	return
}

func RegisterRedisCli(config Config) *redis.Client {
	return initRedis(getRedis(config))
}
