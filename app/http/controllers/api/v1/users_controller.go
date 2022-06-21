package v1

import (
    "GDColumn/app/models/user"
    "GDColumn/app/requests"
    "GDColumn/pkg/auth"
    "GDColumn/pkg/config"
    "GDColumn/pkg/file"
    "GDColumn/pkg/response"
    "GDColumn/pkg/snowflake"
    "github.com/gin-gonic/gin"
)


type UsersController struct {
    BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
    userModel := auth.CurrentUser(c)
    response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context) {
    request := requests.PaginationRequest{}
    if ok := requests.Validate(c, &request, requests.Pagination); !ok {
        return
    }

    data, pager := user.Paginate(c, 10)
    response.JSON(c, gin.H{
        "data":  data,
        "pager": pager,
    })
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

    request := requests.UserUpdateProfileRequest{}
    if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
        return
    }

    currentUser := auth.CurrentUser(c)
    currentUser.NickName = request.NickName
    currentUser.Description = request.Description
    rowsAffected := currentUser.Save()
    if rowsAffected > 0 {
        response.Data(c, currentUser)
    } else {
        response.Abort500(c, "更新失败，请稍后尝试~")
    }
}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {

    request := requests.UserUpdateEmailRequest{}
    if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
        return
    }

    currentUser := auth.CurrentUser(c)
    currentUser.Email = request.Email
    rowsAffected := currentUser.Save()

    if rowsAffected > 0 {
        response.Success(c)
    } else {
        // 失败，显示错误提示
        response.Abort500(c, "更新失败，请稍后尝试~")
    }
}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {

    request := requests.UserUpdatePhoneRequest{}
    if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
        return
    }

    currentUser := auth.CurrentUser(c)
    currentUser.Phone = request.Phone
    rowsAffected := currentUser.Save()

    if rowsAffected > 0 {
        response.Success(c)
    } else {
        response.Abort500(c, "更新失败，请稍后尝试~")
    }
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {

    request := requests.UserUpdatePasswordRequest{}
    if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
        return
    }

    currentUser := auth.CurrentUser(c)
    // 验证原始密码是否正确
    _, err := auth.Attempt(currentUser.NickName, request.Password)
    if err != nil {
        // 失败，显示错误提示
        response.Unauthorized(c, "原密码不正确")
    } else {
        // 更新密码为新密码
        currentUser.Password = request.NewPassword
        currentUser.Save()

        response.Success(c)
    }
}

func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {

    request := requests.UserUpdateAvatarRequest{}
    if ok := requests.Validate(c, &request, requests.UserUpdateAvatar); !ok {
        return
    }
    avatarId,_ := snowflake.GetID()
    currentUser := auth.CurrentUser(c)
    if currentUser.AvatarID != 0 {
        avatar, err := file.SaveUploadAvatar(currentUser.AvatarID,c, request.Avatar)
        if err != nil {
            response.Abort500(c, "上传头像失败，请稍后尝试~")
            return
        }
        currentUser.Avatar.ID = currentUser.AvatarID
        currentUser.Avatar.URL = config.GetString("app.url") + avatar
        currentUser.Save()
        response.Data(c, currentUser)
    }else {
        avatar, err := file.SaveUploadAvatar(avatarId,c, request.Avatar)
        //var oss oss.OssStruct
        //oss.LocalUrl(request.Avatar)
        //oss.Bucket.PutObject("https://bitpig-column.oss-cn-hangzhou.aliyuncs.com/exampledir")
        if err != nil {
            response.Abort500(c, "上传头像失败，请稍后尝试~")
            return
        }

        currentUser.AvatarID =  avatarId
        currentUser.Avatar.ID = currentUser.AvatarID
        currentUser.Avatar.URL = config.GetString("app.url") + avatar
        currentUser.Save()

        response.Data(c, currentUser)
    }
}