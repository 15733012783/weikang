package mysql

import (
	"gorm.io/gorm"
	"time"
)

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

type Goods struct {
	gorm.Model
	GoodsName   string  `gorm:"type:varchar(50)"`
	GoodsPrice  float64 `gorm:"type:decimal(10,2)"`
	GoodsStock  int64   `gorm:"type:int(5)"`
	GoodsType   string  `gorm:"type:varchar(255)"`
	Description string  `gorm:"type:varchar(255)"`
	PageView    int64   `gorm:"type:int(5)"`
}

type Orders struct {
	ID          int
	GoodsID     int
	OrderID     int
	UserID      int
	OrderNumber string
	Status      int8
	TotalPrice  string
	Payment     string
	Address     string
	Quantity    int
	OrderAt     time.Time
}
