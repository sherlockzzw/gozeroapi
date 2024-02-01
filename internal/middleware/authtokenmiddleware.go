package middleware

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"net/http"
)

type AuthTokenMiddleware struct {
	Config config.Config
}

func NewAuthTokenMiddleware(c config.Config) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		Config: c,
	}
}

func (m *AuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers, code, err := result.New(r, m.Config.Auth.AccessExpire)
		if err != nil {
			result.ResultFail(r.Context(), w, code, err)
			return
		}
		checkCode, err := headers.Check()
		if err != nil {
			result.ResultFail(r.Context(), w, checkCode, err)
			return
		}
		next(w, r)
		w.Header()
		token, checkAtCode, err := headers.CheckAccessExpireTime(m.Config.Auth.AccessSecret, m.Config.Auth.AccessExpire)
		if err != nil {
			result.ResultFail(r.Context(), w, checkAtCode, err)
			return
		}
		if token != "" {
			w.Header().Add("token", token)
		}
	}
}
