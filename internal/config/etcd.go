package config

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	configOnce  sync.Once
	RedisClient *redis.Client
	EtcdClient  *clientv3.Client
)

type Server struct{}

func InItWatchEtcd(config Config) {
	wg := new(sync.WaitGroup)
	RegisterEtcd(context.Background(), &config)
	watchRedis(config, wg)
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
					filePath, err = filepath.Abs("etc/apiclipfilm.yaml")
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

func RegisterEtcd(ctx context.Context, config *Config) {
	var err error
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   config.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	logx.WithContext(ctx).Info("successfully connect etcd client")
}

func watchRedis(config Config, wg *sync.WaitGroup) {
	defer wg.Done()
	watchDBChan := EtcdClient.Watch(context.Background(), config.RedisX)
	wg.Add(1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				logx.Errorf("watch redis func panic: %s", p)
			}
		}()
		for {
			consume := <-watchDBChan
			redisX := &RedisX{}
			if err := json.Unmarshal(consume.Events[0].Kv.Value, redisX); err != nil {
				panic(err)
			}
			initRedis(redisX)
			return
		}
	}()
}

func initRedis(redisX *RedisX) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Username: redisX.UserName,
		Addr:     redisX.Addr,
		Password: redisX.PassWord, // 密码
		DB:       redisX.Db,       // 默认数据库
	})
	// 设置连接超时时间

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	logx.Infof("successfully connect redis client")

	return RedisClient
}

func RegisterRedisCli(config Config) *redis.Client {
	return initRedis(getRedis(config))
}
