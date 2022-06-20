package column

import (
    "GDColumn/app/models"
    "GDColumn/pkg/database"
)

type Column struct {
    models.BaseModel
    CID         uint64 `json:"cid"`
    Title       string `json:"title"`
    Description string `json:"description"`
    AvatarID    uint64 `json:"avatar_id"`
    Avatar      Avatar
    Author      uint64 `json:"author"`
    models.CommonTimestampsField
}

type Avatar struct {
    ID  uint64 `json:"_id"`
    URL string `json:"url"`
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