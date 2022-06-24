package column

import (
    "GDColumn/app/models"
    "GDColumn/pkg/database"
)

type Column struct {
    ID          string `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
    Title       string `json:"title,omitempty"`
    Description string `json:"description,omitempty"`
    AvatarID    string `json:"-"`

    Avatar      *Image  `json:"avatar,omitempty"`
    Author      string `json:"author,omitempty"`
    models.CommonTimestampsField
}

type Image struct {
    ID  string `json:"id,omitempty"`
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