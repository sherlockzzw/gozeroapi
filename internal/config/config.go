package config

import "github.com/zeromicro/go-zero/rest"

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
