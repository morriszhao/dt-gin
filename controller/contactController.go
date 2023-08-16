package controller

import (
	"github.com/gin-gonic/gin"
	"morris/im/helper"
	userRequest "morris/im/requests"
	"morris/im/responses"
	"morris/im/services"
)

type ContactController struct {
}

// AddFriend 添加好友
func (ct *ContactController) AddFriend(c *gin.Context) {
	addFriendReq := &userRequest.AddFriendRequest{}
	err := c.ShouldBindJSON(addFriendReq)
	if err != nil {
		helper.RespFail(c, helper.RequestParamsError, err.Error())
		return
	}

	err = services.NewContactServices(c).AddFriend(addFriendReq)
	if err != nil {
		helper.RespFail(c, helper.SystemError, err.Error())
		return
	}

	helper.RespOk(c, helper.Ok, "添加成功")
}

// LoadFriend 好有列表
func (ct *ContactController) LoadFriend(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userFriends, err := services.NewContactServices(c).LoadFriends(int(userId.(int64)))
	if err != nil {
		helper.RespFail(c, helper.SystemError, err.Error())
		return
	}

	helper.RespOk(c, helper.Ok, new(responses.ContactResponse).FriendsResponse(userFriends))
}

func (ct *ContactController) CreateCommunity() {

}

func (ct *ContactController) LoadCommunity() {

}

func (ct *ContactController) JoinCommunity() {

}
