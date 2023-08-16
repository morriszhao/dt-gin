package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"morris/im/models"
	userRequest "morris/im/requests"
)

type contactServices struct {
	ctx *gin.Context
}

func NewContactServices(ctx *gin.Context) *contactServices {
	return &contactServices{
		ctx: ctx,
	}
}

func (cc *contactServices) AddFriend(addFriendReq *userRequest.AddFriendRequest) error {
	userId := GetUserId(cc.ctx)
	dstId := addFriendReq.DstId
	if userId == dstId {
		return errors.New("不能添加自己哦")
	}

	dstUserInfo := new(models.UserModel)
	err := models.GetOneById(dstId, dstUserInfo)
	if err != nil {
		return err
	}

	if dstUserInfo.Id <= 0 {
		return errors.New("用户不存在")
	}

	//验证是否已经添加
	contactInfo := new(models.ContactModel)
	err = contactInfo.GetUserFriendById(userId, dstId)
	if err != nil {
		return err
	}

	if contactInfo.Id > 0 {
		return errors.New("该用户已经添加")
	}

	err = new(models.ContactModel).AddFriend(userId, dstId)
	return err
}

func (cc *contactServices) LoadFriends(userId int) (userFriendList []models.UserModel, err error) {

	friends, err := new(models.ContactModel).GetUserFriends(userId, 1, 100)
	if err != nil {
		return userFriendList, err
	}

	userIds := make([]int, len(friends))
	for index, contactInfo := range friends {
		userIds[index] = contactInfo.Dstobj
	}

	userFriendList, err = new(models.UserModel).GetUserListByIds(userIds)
	return userFriendList, err
}
