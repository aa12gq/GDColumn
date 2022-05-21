package column

import (
    "GDColumn/app/models"
    "GDColumn/pkg/database"
)

type Column struct {
    models.BaseModel

    CID             uint64 `json:"c_id"`
    Author          uint64 `json:"author"`
    Title           string `json:"title"`
    Description     string `json:"description"`
    Avatar        string `json:"avatar,omitempty"`

    models.CommonTimestampsField
}

func (column *Column) Create() {
    database.DB.Create(&column)
}

func (column *Column) Save() (rowsAffected int64) {
    result := database.DB.Save(&column)
    return result.RowsAffected
}

func (column *Column) Delete() (rowsAffected int64) {
    result := database.DB.Delete(&column)
    return result.RowsAffected
}