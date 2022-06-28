package requests

import (
    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type ColumnRequest struct {
    AvatarID    string `valid:"avatarId" json:"avatarId"`
    Title       string `valid:"title" json:"title,omitempty"`
    Description string `valid:"description" json:"description,omitempty"`
}

func ColumnSave(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "avatarId":     []string{"min:6"},
        "title":        []string{"min_cn:2", "max_cn:20"},
        "description": []string{"min_cn:3", "max_cn:255"},
    }
    messages := govalidator.MapData{
        "avatarId":  []string{
            "min:图片id最小为 6 位",
        },
        "title": []string{
            "min_cn:专栏标题需至少 3 个字",
            "max_cn:长度不能超过 20 个字",
        },
        "description": []string{
            "min_cn:描述长度需至少 3 个字",
            "max_cn:描述长度不能超过 255 个字",
        },
    }
    return validate(data, rules, messages)
}