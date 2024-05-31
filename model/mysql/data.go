package mysql

import "gorm.io/gorm"

type Device struct {
	Id       int    //ID
	Name     string //名称
	Model    string //型号
	Brand    int    //品牌ID
	Classify int    //分类ID
}

type DeviceBrand struct {
	Id    int    //ID
	Name  string //名称
	Addr  string //地址
	Phone string //联系方式
}

type DeviceClassify struct {
	Id   int    //ID
	Name string //名称
	Pid  int    //父级ID
}

type User struct {
	Id       int    //ID
	Username string //名称
	Password string //密码
	Email    string //
	Phone    string //电话
}

type Score struct {
	Id     int //ID
	UserId int //用户id
	Score  int //积分
}

type ChatGptHistory struct {
	gorm.Model
	ChatRoomID       int64  //聊天室ID
	SenderID         int64  //用户ID
	Contents         string //内容
	TheTypeOfMessage string //消息类型
}
