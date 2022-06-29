package v1

import (
    "GDColumn/app/models/image"
    "GDColumn/app/requests"
    "GDColumn/pkg/file"
    "GDColumn/pkg/response"
    "GDColumn/pkg/snowflake"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"
)

type ImagesController struct {
    BaseAPIController
}

// Upload 上传文件
// @Summary 上传文件
// @Description 专栏ID
// @Tags file 关于文件上传的一些操作信息
// @Param file formData file true "文件"
// @Success 200 {object} image.Image
// @Router /upload [POST]
func (ctrl *ImagesController) Upload(c *gin.Context) {

    request := requests.ImageRequest{}
    if ok := requests.Validate(c, &request, requests.ImageUpload); !ok {
        return
    }
    avatarId,_ := snowflake.GetID()
    var imageModel image.Image
    var err error
    imageModel.URL,err = file.SaveUploadImage(cast.ToString(avatarId),c,request.Image)
    if err != nil {
        fmt.Println("上传失败:", err)
        response.Abort500(c, "上传头像失败，请稍后尝试~")
        return
    }
    imageModel.ID = cast.ToString(avatarId)
    imageModel.Save()
    response.Data(c, imageModel)
}