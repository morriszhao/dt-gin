package services

import (
	"encoding/json"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"morris/im/helper"
	"morris/im/models"
	"morris/im/requests"
	"time"
)

type userServices struct {
	ctx *gin.Context
}

func NewUserServices(ctx *gin.Context) *userServices {
	return &userServices{ctx: ctx}
}

func (u *userServices) Register(userRegisterReq *userRequest.UserRegisterRequest) (models.UserModel, error) {
	//todo  手机号是否存在
	mobile := userRegisterReq.Mobile
	userModel := models.UserModel{}

	err := userModel.FindUserByMobile(mobile)
	if err != nil {
		return userModel, err
	}

	if userModel.Id > 0 {
		return userModel, errors.New("此用户已存在")
	}

	//todo  存在则返回错误  不存在则添加数据
	err = userModel.Register(userRegisterReq)
	return userModel, err
}

func (u *userServices) Login(userLoginReq *userRequest.UserLoginRequest) (models.UserModel, error) {
	userModel := models.UserModel{}
	err := userModel.FindUserByMobile(userLoginReq.Mobile)
	spew.Dump(err, userLoginReq.Mobile)
	if err != nil {
		return userModel, err
	}

	if userModel.Id <= 0 {
		return userModel, errors.New("用户不存在")
	}

	plainPasswd := helper.Md5(userLoginReq.PassWd)
	if userModel.Passwd != plainPasswd {
		return userModel, errors.New("密码错误")
	}

	//刷新token
	userModel.Token = helper.Md5(helper.RandString(16))
	userModel.Update()

	//token 缓存
	u.SetUserInfoByToken(userModel)

	return userModel, nil
}

func (u *userServices) Search(userSearchRequest *userRequest.UserSearchRequest) (models.UserModel, error) {

	//优先从缓存获取
	userModel, err := u.searchUserByRedis(userSearchRequest)
	if err == nil {
		return userModel, nil
	}

	err = userModel.FindUserByMobile(userSearchRequest.Mobile)
	if err != nil {
		return userModel, err
	}

	if userModel.Id <= 0 {
		return userModel, errors.New("用户不存在!")
	}

	//缓存
	userInfoBytes, _ := json.Marshal(userModel)
	helper.RedisSet(u.redisUserKey(userSearchRequest.Mobile), string(userInfoBytes), time.Second*3600)
	return userModel, err
}

func (u *userServices) searchUserByRedis(userSearchRequest *userRequest.UserSearchRequest) (models.UserModel, error) {
	userInfo := models.UserModel{}
	userInfoJson := helper.RedisGet(u.redisUserKey(userSearchRequest.Mobile))
	if len(userInfoJson) == 0 {
		return userInfo, errors.New("用户不存在")
	}

	err := json.Unmarshal([]byte(userInfoJson), &userInfo)
	if err != nil {
		return userInfo, err
	}

	return userInfo, nil
}

func (u *userServices) SetUserInfoByToken(userInfo models.UserModel) error {

	cacheKey := u.redisUserTokenKey(userInfo.Token)
	userInfoJson, _ := json.Marshal(userInfo)
	helper.RedisSet(cacheKey, string(userInfoJson), time.Second*3600)
	return nil
}

func (u *userServices) GetUserInfoByToken(token string) (userInfo models.UserModel, err error) {

	cacheKey := u.redisUserTokenKey(token)
	userInfoJson := helper.RedisGet(cacheKey)
	if userInfoJson == "" {
		return models.UserModel{}, nil
	}

	err = json.Unmarshal([]byte(userInfoJson), &userInfo)
	return userInfo, err
}

func (u *userServices) redisUserKey(mobile string) string {
	return "user:" + mobile
}

func (u *userServices) redisUserTokenKey(token string) string {
	return "usertoken:" + token
}
