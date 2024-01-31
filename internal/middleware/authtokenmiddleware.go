package middleware

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type AuthTokenMiddleware struct {
	Config      config.Config
	RedisClient *redis.Client
}

func NewAuthTokenMiddleware(c config.Config, redisClient *redis.Client) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		Config:      c,
		RedisClient: redisClient,
	}
}

func (m *AuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers, code, err := result.New(m.RedisClient, &w, r, m.Config.Auth.AccessExpire)
		if err != nil {
			result.ResultFail(r.Context(), w, code, err)
			return
		}
		checkCode, err := headers.Check(r.Header.Get("source"), r.Header.Get("business"))
		if err != nil {
			result.ResultFail(r.Context(), w, checkCode, err)
			return
		}
		next(w, r)
		token, checkAtCode, err := headers.CheckAccessExpireTime(m.Config.Auth.AccessExpire)
		if err != nil {
			result.ResultFail(r.Context(), w, checkAtCode, err)
			return
		}
		if token != "" {
			w.Header().Add("renewToken", token)
		}
	}
}
