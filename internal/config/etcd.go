package config

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

var (
	configOnce  sync.Once
	RedisClient *redis.Client
	etcdClient  *clientv3.Client
)

type Server struct{}

func InItWatchEtcd(config Config) {
	wg := new(sync.WaitGroup)
	RegisterEtcd(context.Background(), &config)
	watchRedis(config, wg)
}

func RegisterEtcd(ctx context.Context, config *Config) {
	var err error
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   config.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	if _, err = etcdClient.Get(ctx, "ping"); err != nil {
		panic("failed connect etcd client")
	}
	logx.WithContext(ctx).Info("successfully connect etcd client")
}

func watchRedis(config Config, wg *sync.WaitGroup) {
	defer wg.Done()
	watchDBChan := etcdClient.Watch(context.Background(), config.RedisX)
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
