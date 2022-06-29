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


// Store 新建文章
// @Summary 新建文章
// @Description 需要标题、内容
// @Tags post 关于文章的一些操作信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param body body requests.PostRequest true "新建文章，需要提供标题和内容"
// @Success 200 {object} post.Post
//@Router /posts [POST]
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
    var r = []rune(request.Content)
    postModel := &post.Post{
        ID:         cast.ToString(id),
        Title:      request.Title,
        Content:    request.Content,
        Excerpt:    string(r[:20]),
        ImageID:    request.ColumnID,
        AuthorID:   request.AuthorID,
        ColumnID:   request.ColumnID,
        Author:     &userModel,
        Image:      img,
    }
    postModel.Create()
    if postModel.ID != "" {
        response.Data(c, postModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}

// Update 更新单个文章信息
// @Summary 发送请求，更新单个文章信息
// @Tags post 关于文章的一些操作信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer true "文章ID"
// @Param body body requests.PostRequest true "待更新文章数据"
// @Success 200 {object} post.Post
//@Router /posts/{id} [PUT]
func (ctrl *PostsController) Update(c *gin.Context) {

    postModel := post.Get(c.Param("id"))
    if postModel.ID == "" {
        response.Abort404(c)
        return
    }
    request := requests.PostRequest{}
    if ok := requests.Validate(c, &request, requests.PostSave); !ok {
        return
    }
    if ok := policies.CanModifyPost(c, postModel.AuthorID); !ok {
        response.Abort403(c)
        return
    }
    rowsAffected := postModel.Updates(request.ImageID,request.Title,request.Content)
    postModel = post.Get(c.Param("id"))
    if rowsAffected > 0 {
        response.Data(c, postModel)
    } else {
        response.Abort500(c, "更新失败，请稍后尝试~")
    }
}

// Delete 删除个文章
// @Summary 发送请求，删除单个文章
// @Tags post 关于文章的一些操作信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path integer true "文章ID"
// @Success 200 {object} post.Post
//@Router /posts/{id} [DELETE]
func (ctrl *PostsController) Delete(c *gin.Context) {

    postModel := post.Get(c.Param("id"))
    if postModel.ID == "" {
        response.Abort404(c)
        return
    }
    if ok := policies.CanModifyPost(c, postModel.AuthorID); !ok {
        response.Abort403(c)
        return
    }
    rowsAffected := postModel.Delete()
    if rowsAffected > 0 {
        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "删除成功！",
            "data":postModel,
        })
        return
    }
    response.Abort500(c, "删除失败，请稍后尝试~")
}

// Show 获取单个文章信息
// @Summary 发送请求获取单个文章信息
// @Description 文章ID
// @Tags post 关于文章的一些操作信息
// @Param id path integer true "文章ID"
// @Success 200 {object} post.Post
//@Router /posts/{id} [GET]
func (ctrl *PostsController) Show(c *gin.Context) {

    postModel := post.Get(c.Param("id"))
    switch {
    case postModel.ID == "":
        response.Abort404(c)
        break
    case postModel.Author.AvatarID != "":
        imgModel := image.Get(postModel.Author.AvatarID)
        img := &user.Image{
            ID:imgModel.ID,
            URL:imgModel.URL,
        }
        postModel.Author.Avatar = img
    }
    response.Data(c, postModel)
}

// Index 获取属于专栏的文章信息
// @Summary 发送请求获取属于专栏的文章信息
// @Description 专栏ID
// @Tags post 关于文章的一些操作信息
// @Param id path integer true "专栏ID"
// @Success 200 {object} post.Post
// @Router /columns/{id}/posts [GET]
func (ctrl *PostsController) Index(c *gin.Context) {

    postModel := post.GetAll(c.Param("id"))
    for i := 0;i < len(postModel); i++ {
       if postModel[i].ImageID != ""{
           imgModel := image.Get(cast.ToString(postModel[i].ImageID))
           image := &post.Image{
               ID:  imgModel.ID,
               URL: imgModel.URL,
           }
           postModel[i].Image = image
       }
    }
    data,pager :=  post.Paginate(c,5,c.Param("id"))
    for i := 0; i<len(data) ;i++  {
        if data[i].Author.AvatarID != ""{
            imgModel := image.Get(cast.ToString(data[i].Author.AvatarID))
            image := &user.Image{
                ID:  imgModel.ID,
                URL: imgModel.URL,
            }
            data[i].Author.Avatar = image
        }
    }
     response.JSON(c, gin.H{
         "data":  data,
         "pager": pager,
     })
}