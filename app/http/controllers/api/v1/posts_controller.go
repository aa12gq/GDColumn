package v1

import (
    "GDColumn/app/models/image"
    "GDColumn/app/models/post"
    "GDColumn/app/models/user"
    "GDColumn/app/policies"
    "GDColumn/app/requests"
    "GDColumn/pkg/response"
    "GDColumn/pkg/snowflake"
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"
    "net/http"
)

type PostsController struct {
    BaseAPIController
}

func (ctrl *PostsController) Store(c *gin.Context) {

    request := requests.PostRequest{}
    if ok := requests.Validate(c, &request, requests.PostSave); !ok {
        return
    }

    userModel := user.Get(request.AuthorID)
    imgModel := image.Get(request.ImageID)
    img := &post.Image{
       ID:imgModel.ID,
       URL:imgModel.URL,
    }
    id,_:=snowflake.GetID()
    postModel := &post.Post{
        ID:         id,
        Title:      request.Title,
        Content:    request.Content,
        Excerpt:    request.Content,
        ImageID:    cast.ToUint64(request.ColumnID),
        AuthorID:   cast.ToUint64(request.AuthorID),
        ColumnID:   cast.ToUint64(request.ColumnID),
        Author:     &userModel,
        Image:      img,
    }
    if ok := policies.CanModifyPost(c, postModel.AuthorID); !ok {
        response.Abort403(c)
        return
    }
    postModel.Create()
    if postModel.ID > 0 {
        response.Data(c, postModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}

func (ctrl *PostsController) Update(c *gin.Context) {

    postModel := post.Get(c.Param("id"))
    if postModel.ID == 0 {
        response.Abort404(c)
        return
    }

    if ok := policies.CanModifyPost(c, postModel.AuthorID); !ok {
        response.Abort403(c)
        return
    }

    request := requests.PostRequest{}
    if ok := requests.Validate(c, &request, requests.PostSave); !ok {
        return
    }

    imgModel := image.Get(request.ImageID)
    image := &post.Image{
        ID:  imgModel.ID,
        URL:imgModel.URL,
    }

    postModel.Title = request.Title
    postModel.Content = request.Content
    postModel.ImageID = cast.ToUint64(request.ImageID)
    postModel.Image = image

    rowsAffected := postModel.Save()
    if rowsAffected > 0 {
        response.Data(c, postModel)
    } else {
        response.Abort500(c, "更新失败，请稍后尝试~")
    }
}

func (ctrl *PostsController) Delete(c *gin.Context) {

    postModel := post.Get(c.Param("id"))
    if postModel.ID == 0 {
        response.Abort404(c)
        return
    }
    if ok := policies.CanModifyPost(c, postModel.AuthorID); !ok {
        response.Abort403(c)
        return
    }
    rowsAffected := postModel.Delete()
    if rowsAffected > 0 {
        response.Success(c)
        return
    }
    response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *PostsController) Show(c *gin.Context) {

    postModel := post.Get(c.Param("id"))
    if postModel.ID == 0 {
        response.Abort404(c)
        return
    }
    response.Data(c, postModel)
}

func (ctrl *PostsController) Index(c *gin.Context) {

    postModel := post.GetAll(c.Param("id"))
    for i := 0;i < len(postModel); i++ {
       if postModel[i].ImageID != 0{
           imgModel := image.Get(cast.ToString(postModel[i].ImageID))
           image := &post.Image{
               ID:  imgModel.ID,
               URL: imgModel.URL,
           }
           postModel[i].Image = image
       }
    }

    c.JSON(http.StatusOK,gin.H{
        "count":len(postModel) + 1,
        "list": postModel,
        "pageSize":len(postModel),
        "currentPage":1,
    })
}