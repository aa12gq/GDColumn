package v1

import (
    "GDColumn/app/models/column"
    "GDColumn/app/models/image"
    "GDColumn/app/requests"
    "GDColumn/pkg/auth"
    "GDColumn/pkg/response"
    "GDColumn/pkg/snowflake"
    "github.com/spf13/cast"

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
    columnModel := column.Column{
       ID:          cast.ToString(Cid),
       Author:      author,
       Title:       request.Title,
       Description: request.Description,
    }
    columnModel.Create()
    if columnModel.ID != "" {
        response.Created(c, columnModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}

// Update 更新个人专栏信息
// @Summary 发送请求，更新个人专栏信息
// @Tags column 关于专栏的一些操作信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param body body requests.ColumnRequest true "待更新专栏数据"
// @Success 200 {object} column.Column
//@Router /columns/ [PUT]
func (ctrl *ColumnsController) Update(c *gin.Context) {

    request := requests.ColumnRequest{}
    if ok := requests.Validate(c, &request, requests.ColumnSave); !ok {
        return
    }
    currentUser := auth.CurrentUser(c)
    columnModel := column.Get(currentUser.ColumnID)
    rowsAffected := columnModel.Updates(request.AvatarID,request.Title,request.Description)
    columnModel = column.Get(currentUser.ColumnID)
    if rowsAffected > 0 {
        response.Data(c, columnModel)
    } else {
        response.Abort500(c)
    }
}

// ShowColumn 发送请求-获取一个专栏详情
// @Summary 发送请求，获取一个专栏详情
// @Tags column 关于专栏的一些操作信息
// @Param id path integer true "专栏ID"
// @Success 200 {object} column.Column
//@Router /columns/{id} [GET]
func (ctrl *ColumnsController) ShowColumn(c *gin.Context) {

    columnModel := column.Get(c.Param("id"))
    switch {
    case columnModel.ID == "":
        response.Abort404(c)
        break
    case columnModel.AvatarID != "":
        imgModel := image.Get(columnModel.AvatarID)
        avatar := &column.Image{
            ID:  imgModel.ID,
            URL: imgModel.URL,
        }
        columnModel.Avatar = avatar
    }

    if columnModel.ID > "" {
        response.Data(c, columnModel)
    } else {
        response.Abort500(c)
    }
}

// Index 获取专栏列表
// @Summary 发送请求，获得专栏列表
// @Tags column 关于专栏的一些操作信息
// @Accept application/json
// @Produce application/json
// @Param object query requests.PaginationRequest false "查询参数"
// @Success 200 {object} column.Column
//@Router /columns [GET]
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

func (ctrl *ColumnsController) Delete(c *gin.Context) {

    columnModel := column.Get(c.Param("id"))
    if columnModel.ID == "" {
        response.Abort404(c)
        return
    }
    author := auth.CurrentUID(c)
    if author == columnModel.Author {
        rowsAffected := columnModel.Delete()
        if rowsAffected > 0 {
            response.Success(c)
            return
        }
    }else {
        response.Abort500(c, "删除失败，请稍后尝试~")
    }
}