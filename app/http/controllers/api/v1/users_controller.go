package v1

import (
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