package user

import (
	"awesomeProject/goweb/pkg/logger"
	"awesomeProject/goweb/pkg/model"
	"awesomeProject/goweb/pkg/types"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Get 通过 ID 获取文章
func Get(idstr string) (User, error) {
	var user User
	id := types.StringToInt(idstr)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

// 通过邮箱获取用户
func GetByEmail(email string) (User, error) {

	var user User

	if err := model.DB.Where("email", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(password string) bool {
	return u.Password == password
}
