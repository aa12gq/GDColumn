package requests

import (
    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type PostRequest struct {
    Title      string `json:"title,omitempty" valid:"title"`
    Content    string `json:"content,omitempty" valid:"content"`
    ImageID    string `json:"image_id,omitempty" valid:"image_id"`
    ColumnID     string `json:"column_id,omitempty" valid:"column_id"`
    AuthorID     string `json:"author_id,omitempty" valid:"author_id"`
}

func PostSave(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "title":       []string{"required", "min_cn:3", "max_cn:40"},
        "content":     []string{"required", "min_cn:5", "max_cn:50000"},
        "image_id":       []string{ "min_cn:10", "max_cn:50000"},
        "column_id":   []string{"exists:columns,id"},
        "author_id" : []string{"exists:users,id"},
}
    messages := govalidator.MapData{
        "title": []string{
            "required:帖子标题为必填项",
            "min_cn:标题长度需大于 3",
            "max_cn:标题长度需小于 40",
        },
        "image_id": []string{
            "exists:图片未找到",
        },
        "content": []string{
            "required:文章内容为必填项",
            "min_cn:长度需大于 5",
        },
        "column_id": []string{
            "exists:文章的专栏未找到",
        },
        "author_id": []string{
            "exists:作者未找到",
        },
    }
    return validate(data, rules, messages)
}