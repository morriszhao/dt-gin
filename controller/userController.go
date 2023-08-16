package controller

import (
	"github.com/gin-gonic/gin"
	"morris/im/helper"
	"morris/im/requests"
	"morris/im/responses"
	"morris/im/services"
)

type UserController struct {
}

// Login  登录
func (u *UserController) Login(c *gin.Context) {
	userLoginRequest := &userRequest.UserLoginRequest{}
	err := c.ShouldBindJSON(userLoginRequest)
	if err != nil {
		helper.RespFail(c, helper.RequestParamsError, err.Error())
		return
	}

	userInfo, err := services.NewUserServices(c).Login(userLoginRequest)
	if err != nil {
		helper.RespFail(c, helper.SystemError, err.Error())
		return
	}

	helper.RespOk(c, helper.Ok, new(responses.UserResponse).LoginResponse(userInfo))
}

// Register 注册
func (u *UserController) Register(c *gin.Context) {

	//结构体验证
	userRegisterRequest := &userRequest.UserRegisterRequest{}

	//ShouldBindJSON 与 BindJson 区别； BindJson 验证未通过会写 400 状态码
	err := c.ShouldBindJSON(userRegisterRequest)
	if err != nil {
		helper.RespFail(c, helper.RequestParamsError, err.Error())
		return

	}

	userInfo, err := services.NewUserServices(c).Register(userRegisterRequest)
	if err != nil {
		helper.RespFail(c, helper.SystemError, err.Error())
		return
	}

	helper.RespOk(c, helper.Ok, new(responses.UserResponse).RegisterResponse(userInfo))

}

// Search 搜索
func (u *UserController) Search(c *gin.Context) {
	userSearchRequest := &userRequest.UserSearchRequest{}
	err := c.ShouldBindQuery(userSearchRequest)
	if err != nil {
		helper.RespFail(c, helper.RequestParamsError, err.Error())
		return
	}

	userInfo, err := services.NewUserServices(c).Search(userSearchRequest)
	if err != nil {
		helper.RespFail(c, helper.SystemError, err.Error())
		return
	}

	helper.RespOk(c, helper.Ok, new(responses.UserResponse).SearchResponse(userInfo))
}
