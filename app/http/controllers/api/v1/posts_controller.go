package v1

import (
    "GDColumn/app/models/post"
    "GDColumn/app/requests"
    "GDColumn/pkg/auth"
    "GDColumn/pkg/response"

    "github.com/gin-gonic/gin"
    "fmt"
)

type PostsController struct {
    BaseAPIController
}

func (ctrl *PostsController) Store(c *gin.Context) {

    request := requests.PostRequest{}
    if ok := requests.Validate(c, &request, requests.PostSave); !ok {
        return
    }
    fmt.Println("通过")


    postModel := post.Post{
        Title:      request.Title,
        Content:     request.Content,
        Excerpt:    request.Excerpt,
        ColumnID:    request.ColumnID,
        UserID:     auth.CurrentUID(c),
    }
    postModel.Create()
    if postModel.ID > 0 {
        response.Created(c, postModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}