package requests

import (
    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type PostRequest struct {
    Title      string `json:"title,omitempty"    valid:"title"`
    Content    string `json:"content,omitempty"  valid:"content"`
    ImageID    string `json:"image,omitempty"    valid:"image"`
    ColumnID   string `json:"columnId,omitempty" valid:"columnId"`
    AuthorID   string `json:"authorId,omitempty" valid:"authorId"`
}

func PostSave(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "title":       []string{ "min_cn:3", "max_cn:40"},
        "content":     []string{ "min_cn:5", "max_cn:50000"},
        "image":       []string{ "min_cn:10", "max_cn:50000"},
        "columnId":    []string{"exists:columns,id"},
        "author_id" :  []string{"exists:users,id"},
}
    messages := govalidator.MapData{
        "title": []string{
            "min_cn:标题长度需大于 3",
            "max_cn:标题长度需小于 40",
        },
        "image": []string{
            "exists:图片未找到",
        },
        "content": []string{
            "min_cn:长度需大于 5",
        },
        "columnId": []string{
            "exists:文章的专栏未找到",
        },
        "authorId": []string{
            "exists:作者未找到",
        },
    }
    return validate(data, rules, messages)
}