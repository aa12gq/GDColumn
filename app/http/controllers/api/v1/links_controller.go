package v1

import (
    "GDColumn/app/models/link"
    "GDColumn/pkg/response"

    "github.com/gin-gonic/gin"
)

type LinksController struct {
    BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
    response.Data(c, link.AllCached())
}