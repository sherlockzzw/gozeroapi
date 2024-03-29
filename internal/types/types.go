// Code generated by goctl. DO NOT EDIT.
package types

type GetUserListReq struct {
	Page     int64 `form:"page,default=1,range=(0:]"`      // 分页，默认第一页
	PageSize int64 `form:"pageSize,default=10,range=(0:]"` // 分页查询条数
}

type GetUserListRes struct {
	List []UserList `json:"list"` // list
}

type UserList struct {
	Id         int64  `form:"id"`         // id
	Uuid       string `form:"uuid"`       // uuid
	Mobile     string `form:"mobile"`     // 手机号
	Email      string `form:"email"`      //邮箱
	State      int64  `form:"state"`      //状态
	Created_at int64  `form:"created_at"` //创建时间
	Updated_at int64  `form:"updated_at"` //更新时间
}
