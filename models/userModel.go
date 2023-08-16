package models

import (
	"morris/im/helper"
	userRequest "morris/im/requests"
	"time"
)

const (
	OnlineTrue  = 1
	OnlineFalse = 0
)

type UserModel struct {
	//必须要加  xorm:"extends" tag、 否则会被当成字段进行映射
	BaseModel `xorm:"extends"`
	//用户ID
	Mobile   string `xorm:"varchar(20)" json:"mobile"`
	Passwd   string `xorm:"varchar(40)" json:"passwd,omitempty"`
	Avatar   string `xorm:"varchar(150)" json:"avatar"`
	Sex      string `xorm:"varchar(2)" json:"sex"`
	Nickname string `xorm:"varchar(20)" json:"nickname"`
	Salt     string `xorm:"varchar(10)" json:"salt"`
	Online   int    `xorm:"int(10)" json:"online"`
	Token    string `xorm:"varchar(40)" json:"token"`
	Memo     string `xorm:"varchar(140)" json:"memo"`
}

//TableName  返回表名
func (u *UserModel) TableName() string {
	return "user"
}

// Register 注册用户
func (u *UserModel) Register(userRequest *userRequest.UserRegisterRequest) error {
	u.Mobile = userRequest.Mobile
	u.Passwd = helper.Md5(userRequest.PassWd)
	u.Avatar = userRequest.Avatar
	u.Sex = userRequest.Sex
	u.Nickname = userRequest.Nickname
	u.Salt = helper.RandString(4)
	u.Online = OnlineFalse
	u.Token = helper.Md5(helper.RandString(16))
	u.Createat = time.Now()

	_, err := DbEngin.InsertOne(u)
	return err
}

// FindUserByMobile 根据手机号 查找用户
func (u *UserModel) FindUserByMobile(mobile string) error {
	_, err := DbEngin.Where("mobile=?", mobile).Get(u)
	return err
}

func (u *UserModel) Update() error {
	_, err := DbEngin.Where("id=?", u.Id).Update(u)
	return err
}

func (u *UserModel) GetUserListByIds(userIds []int) ([]UserModel, error) {
	userModelList := make([]UserModel, 0)
	err := DbEngin.In("id", userIds).Find(&userModelList)
	return userModelList, err
}
