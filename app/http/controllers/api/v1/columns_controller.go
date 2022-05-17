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

func (ctrl *ColumnsController) Update(c *gin.Context) {

    // 验证 url 参数 id 是否正确
    columnModel := column.Get(c.Param("id"))
    if columnModel.CID == 0 {
        response.Abort404(c)
        return
    }


    request := requests.ColumnRequest{}
    if ok := requests.Validate(c, &request, requests.ColumnSave); !ok {
        return
    }

    // 保存数据
    columnModel.Title = request.Title
    columnModel.Description = request.Description
    rowsAffected := columnModel.Save()

    if rowsAffected > 0 {
        response.Data(c, columnModel)
    } else {
        response.Abort500(c)
    }
}

func (ctrl *ColumnsController) Index(c *gin.Context) {

    //request := requests.PaginationRequest{}
    //if ok := requests.Validate(c, &request, requests.Pagination); !ok {
    //    return
    //}

    data, pager := column.Paginate(c, 5)
    response.JSON(c, gin.H{
        "data":  data,
        "pager": pager,
    })
}