package responses

import "morris/im/models"

type UserResponse struct {
}

func (u *UserResponse) LoginResponse(data models.UserModel) interface{} {

	data.Id = 0
	data.Passwd = ""
	data.Salt = ""

	return data
}

func (u *UserResponse) RegisterResponse(data models.UserModel) interface{} {
	data.Id = 0
	data.Passwd = ""
	data.Salt = ""

	return data
}

func (u *UserResponse) SearchResponse(data models.UserModel) interface{} {
	data.Id = 0
	data.Passwd = ""
	data.Salt = ""

	return data
}
