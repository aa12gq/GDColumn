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
    links := link.All()
    response.Data(c, links)
}