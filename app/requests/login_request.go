package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func LoginByEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "max:30", "email"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	errs := validate(data, rules, messages)
	return errs
}