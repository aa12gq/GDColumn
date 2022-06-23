package v1

import (
    "GDColumn/app/models/post"
    "GDColumn/app/models/image"
    "GDColumn/app/models/user"
    "GDColumn/app/requests"
    "GDColumn/pkg/response"
    "GDColumn/pkg/snowflake"
    "strconv"

    "fmt"
    "github.com/gin-gonic/gin"
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


    userModel := user.Get(request.UserID)
    avatarModel := image.Get(strconv.FormatUint(userModel.AvatarID,10))
    userModel.Avatar.ID = avatarModel.ID
    userModel.Avatar.URL = avatarModel.URL

    imgModel := image.Get(request.ImageID)
    img := &post.Image{
        ID:imgModel.ID,
        URL:imgModel.URL,
    }

    id,_:=snowflake.GetID()
    postModel := &post.Post{
     ID:         id,
     Title:      request.Title,
     Content:     request.Content,
     Excerpt:    request.Content,
     ColumnID:    request.ColumnID,
     UserID:     request.UserID,
     ImageID:    request.ImageID,
    }

    post2Model := &post.Post2{
        ID:                    id,
        Title:                 request.Title,
        Content:               request.Content,
        Excerpt:               request.Content,
        ImageID:               request.ImageID,
        UserID:                request.UserID,
        ColumnID:              request.ColumnID,
        User:                  userModel,
        Image:                 *img,
    }
    postModel.Create()
    if postModel.ID > 0 {
        response.Data(c, post2Model)
    } else {
        fmt.Println("error")
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}