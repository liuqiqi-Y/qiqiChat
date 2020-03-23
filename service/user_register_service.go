package service

import (
	"qiqiChat/model"
	"qiqiChat/serializer"
	"qiqiChat/util"
	"time"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	row := model.DB.QueryRow(`SELECT COUNT(*) FROM user WHERE user_name = ? AND status = 1`, service.UserName)
	row.Scan(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		UserName:  service.UserName,
		Status:    model.Active,
		CreatedAt: time.Now(),
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	stmt, err := model.DB.Prepare(`INSERT INTO user(user_name, password_digest, status) VALUES(?,?,?)`)
	if err != nil {
		util.Err.Println("Failed to prepare sql statement: ", err.Error())
		return serializer.Err(serializer.CodeDBError, "注册失败", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.UserName, user.PasswordDigest, user.Status)
	if err != nil {
		util.Err.Println("Failed to insert data: ", err.Error())
		return serializer.Err(serializer.CodeDBError, "注册失败", err)
	}
	_ = model.DB.QueryRow(`SELECT id FROM user WHERE user_name = ?`, user.UserName).Scan(&user.ID)
	return serializer.BuildUserResponse(user)
}
