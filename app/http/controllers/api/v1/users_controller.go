package v1

import (
    "GDColumn/pkg/auth"
    "GDColumn/pkg/response"
    "GDColumn/app/models/user"

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
    data := user.All()
    response.Data(c, data)
}