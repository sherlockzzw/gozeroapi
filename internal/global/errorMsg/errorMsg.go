package errorMsg

import "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"

const (
	ForbiddenWordError result.BusinessCode = 12001
)

func New() {
	server := result.NewErrorServer()
	server.Add(ForbiddenWordError, "输入文案包含违禁词", "user not")
}
