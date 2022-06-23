package column

import (
    "GDColumn/app/models"
    "GDColumn/pkg/database"
)

type Column struct {
    ID          uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
    Title       string `json:"title,omitempty"`
    Description string `json:"description,omitempty"`
    AvatarID    uint64 `json:"-"`

    Avatar      *Image  `json:"avatar,omitempty"`
    Author      uint64 `json:"author,omitempty"`
    models.CommonTimestampsField
}

type Image struct {
    ID  uint64 `json:"id,omitempty"`
    URL string `json:"url,omitempty"`
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