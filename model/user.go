package model

import (
	"qiqiChat/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	err := DB.QueryRow(`SELECT id, created_at, deleted_at, user_name, password_digest, nick_name, status, avatar FROM users WHERE id = ?`, ID).Scan(&user.ID, &user.CreatedAt, &user.DeletedAt, &user.UserName, &user.PasswordDigest, &user.Nickname, &user.Status, &user.Avatar)
	if err != nil {
		util.Err.Println("Faile to User information: ", err.Error())
	}
	return user, err
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
