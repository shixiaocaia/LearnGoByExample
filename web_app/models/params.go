package models

// 定义请求的参数结构体

// tag错误, 会导致无法识别
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}
