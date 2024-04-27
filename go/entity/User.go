package entity

// 和我想的一样，可以正常接送请求，但是得不到信息，这是因为估计他表没有找到
// 在mybatis中也有类似的问题

// TableName 定义数据库表名
func (User) TableName() string {
	return "test.user"
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Password string `json:"password"`
}
