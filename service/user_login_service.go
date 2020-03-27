package service

import (
	"qiqiChat/model"
	"qiqiChat/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName       string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password       string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Identification string `form:"identification" json:"identification" binding:"required"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	err := model.DB.QueryRow("SELECT `user_name`, `password_digest`, `id`, `status`, `created_at`, `identification` FROM user WHERE user_name = ? AND `status` = 1 AND `identification` = ?", service.UserName, service.Identification).Scan(&user.UserName,
		&user.PasswordDigest, &user.ID, &user.Status, &user.CreatedAt, &user.Identification)
	if err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
