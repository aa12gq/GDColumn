package auth

import (
	"errors"
	"GDColumn/app/models/user"
)

// Attempt 尝试登录
func Attempt(email, password string)(user.User,error){
	userModel := user.GetByMulti(email)
	if userModel.ID == 0{
		return user.User{},errors.New("用户不存在")
	}
	if !userModel.ComparePassword(password){
		return user.User{},errors.New("密码错误")
	}
	return userModel,nil
}

func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}