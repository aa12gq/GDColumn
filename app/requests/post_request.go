package requests

import (
    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type PostRequest struct {
    Title      string `json:"title,omitempty" valid:"title"`
    Content    string `json:"content,omitempty" valid:"content"`
    Excerpt    string `json:"excerpt,omitempty" valid:"excerpt"`
    Image      string `json:"image,omitempty" valid:"body"`
    ColumnID   string `json:"column_id,omitempty" valid:"column_id"`
    Author     string `json:"author,omitempty" valid:"author"`
}

func PostSave(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "title":       []string{"required", "min_cn:3", "max_cn:40"},
        "content":     []string{"required", "min_cn:5", "max_cn:50000"},
        "excerpt":     []string{"required", "min_cn:5", "max_cn:50000"},
        "Image":       []string{"required", "min_cn:10", "max_cn:50000"},
        "column_id":   []string{"required", "exists:column,id"},
        "author" : []string{"required", "exists:author,id"},
}
    messages := govalidator.MapData{
        "title": []string{
            "required:帖子标题为必填项",
            "min_cn:标题长度需大于 3",
            "max_cn:标题长度需小于 40",
        },
        "image": []string{
            "required:图片内容为必填项",
            "min_cn:长度需大于 10",
        },
        "content": []string{
            "required:文章内容为必填项",
            "min_cn:长度需大于 5",
        },
        "excerpt": []string{
            "required:文章内容为必填项",
            "min_cn:长度需大于 5",
        },
        "column_id": []string{
            "required:文章的专栏为必填项",
            "exists:文章的专栏未找到",
        },
        "author": []string{
            "required:作者id为必填项",
            "exists:作者未找到",
        },
    }
    return validate(data, rules, messages)
}