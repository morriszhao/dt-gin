package userRequest

type UserRegisterRequest struct {
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	PassWd   string `json:"pass_word"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex" binding:"required"`
}

type UserLoginRequest struct {
	Mobile string `json:"mobile" binding:"required"`
	PassWd string `json:"pass_word" binding:"required"`
}

type UserSearchRequest struct {
	Mobile string `form:"mobile" binding:"required"`
}
