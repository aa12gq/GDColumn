package v1

import (
    "GDColumn/app/models/user"
    "GDColumn/app/requests"
    "GDColumn/pkg/auth"
    "GDColumn/pkg/response"
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