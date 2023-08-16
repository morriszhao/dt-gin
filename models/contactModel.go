package models

import (
	"time"
)

const (
	CatePrivateChat = 1
	CatePublicChat  = 2
)

type ContactModel struct {
	BaseModel `xorm:"extends"`
	//自增id
	Ownerid int    `xorm:"bigint(20)"json:"ownerid"` // 记录是谁的
	Dstobj  int    `xorm:"bigint(20)" json:"dstobj"` // 对端信息
	Cate    int    `xorm:"int(11)"json:"cate"`       // 什么类型   1:单聊  2:群聊
	Memo    string `xorm:"varchar(120)" json:"memo"` // 备注
}

//TableName  返回表名
func (cc *ContactModel) TableName() string {
	return "contact"
}

func (cc *ContactModel) AddFriend(userId, dstId int) error {

	//事务添加。
	session := DbEngin.NewSession()
	session.Begin()

	//插入自己的
	_, err1 := session.InsertOne(&ContactModel{
		Ownerid: userId,
		Dstobj:  dstId,
		Cate:    CatePrivateChat,
		Memo:    "",
		BaseModel: BaseModel{
			Createat: time.Now(),
		},
	})

	//插入对方的
	_, err2 := session.InsertOne(&ContactModel{
		Ownerid: dstId,
		Dstobj:  userId,
		Cate:    CatePrivateChat,
		Memo:    "",
		BaseModel: BaseModel{
			Createat: time.Now(),
		},
	})

	//提交事务 货物rollback
	if err1 == nil && err2 == nil {
		session.Commit()
		return nil
	}

	session.Rollback()
	if err1 != nil {
		return err1
	}
	return err2
}

func (cc *ContactModel) GetUserFriendById(userId int, dstId int) error {
	_, err := DbEngin.Where("ownerid=? and dstobj=?", userId, dstId).Get(cc)
	return err
}

func (cc *ContactModel) GetUserFriends(userId int, page int, pageSize int) ([]ContactModel, error) {

	//计算分页
	start := (page - 1) * pageSize

	var contactList []ContactModel
	err := DbEngin.Where("ownerid=?", userId).Limit(pageSize, start).Find(&contactList)
	return contactList, err
}
