package v1

import (
    "GDColumn/app/requests"
    "GDColumn/app/models/image"
    "GDColumn/pkg/file"
    "GDColumn/pkg/response"
    "GDColumn/pkg/snowflake"
    "github.com/gin-gonic/gin"
    "fmt"
    "github.com/spf13/cast"
    "time"
)

type ImagesController struct {
    BaseAPIController
}

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
    imageModel.ID = avatarId
    imageModel.CreatedAt = time.Now()
    imageModel.UpdatedAt = time.Now()
    imageModel.Save()
    response.Data(c, imageModel)
}