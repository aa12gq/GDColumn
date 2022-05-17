package column

import (
     "GDColumn/pkg/app"
     "GDColumn/pkg/database"
     "GDColumn/pkg/paginator"

     "github.com/gin-gonic/gin"
)

func Get(idstr string) (column Column) {
    database.DB.Where("c_id", idstr).First(&column)
    return
}

func GetBy(field, value string) (column Column) {
    database.DB.Where("? = ?", field, value).First(&column)
    return
}

func All() (columns []Column) {
    database.DB.Find(&columns)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Column{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (columns []Column, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Column{}),
        &columns,
        app.V1URL(database.TableName(&Column{})),
        perPage,
    )
    return
}