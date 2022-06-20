package requests

import (
    "GDColumn/pkg/auth"
    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type UserUpdateProfileRequest struct {
    NickName          string `valid:"nick_name" json:"nick_name"`
    Description      string `valid:"description" json:"description"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {

    // 查询用户名重复时，过滤掉当前用户 ID
    uid := auth.CurrentUID(c)
    rules := govalidator.MapData{
        "nick_name":          []string{"required", "alpha_num", "between:3,20", "not_exists:users,nick_name," + uid},
        "description": []string{"min_cn:4", "max_cn:240"},
    }

    messages := govalidator.MapData{
        "nick_name": []string{
            "required:用户名为必填项",
            "alpha_num:用户名格式错误，只允许数字和英文",
            "between:用户名长度需在 3~20 之间",
            "not_exists:用户名已被占用",
        },
        "description": []string{
            "min_cn:描述长度需至少 4 个字",
            "max_cn:描述长度不能超过 240 个字",
        },
    }
    return validate(data, rules, messages)
}

type UserUpdateEmailRequest struct {
    Email      string `json:"email,omitempty" valid:"email"`
    VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {

    currentUser := auth.CurrentUser(c)
    rules := govalidator.MapData{
        "email": []string{
            "required", "min:4",
            "max:30",
            "email",
            "not_exists:users,email," + currentUser.GetStringID(),
            "not_in:" + currentUser.Email,
        },
        "verify_code": []string{"required", "digits:6"},
    }
    messages := govalidator.MapData{
        "email": []string{
            "required:Email 为必填项",
            "min:Email 长度需大于 4",
            "max:Email 长度需小于 30",
            "email:Email 格式不正确，请提供有效的邮箱地址",
            "not_exists:Email 已被占用",
            "not_in:新的 Email 与老 Email 一致",
        },
        "verify_code": []string{
            "required:验证码答案必填",
            "digits:验证码长度必须为 6 位的数字",
        },
    }

    return validate(data, rules, messages)
}