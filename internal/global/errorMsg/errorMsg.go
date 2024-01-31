package errorMsg

import "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"

const (
	UserNot result.BusinessCode = 1000
)

func New() {
	server := result.NewErrorServer()
	server.Add(UserNot, "用户不存在", "user not")
}
