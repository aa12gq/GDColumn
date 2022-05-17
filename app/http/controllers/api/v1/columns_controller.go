package v1

import (
    "GDColumn/app/models/column"
    "GDColumn/app/requests"
    "GDColumn/pkg/auth"
    "GDColumn/pkg/response"
    "GDColumn/pkg/snowflake"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ColumnsController struct {
    BaseAPIController
}


func (ctrl *ColumnsController) Store(c *gin.Context) {

    request := requests.ColumnRequest{}
    if ok := requests.Validate(c, &request, requests.ColumnSave); !ok {
        return
    }
    Cid,_ := snowflake.GetID()
    author := auth.CurrentUID(c)
    Aid,_ := strconv.ParseUint(author, 10, 64)
    columnModel := column.Column{
       CID:         Cid,
       Author:      Aid,
       Title:       request.Title,
       Description: request.Description,
    }
    columnModel.Create()
    if columnModel.CID > 0 {
        response.Created(c, columnModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}
