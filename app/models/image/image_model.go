package image

import (
    "GDColumn/app/models"
    "GDColumn/pkg/database"
)

type Image struct {
    ID       uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
    URL      string `json:"url"`

    models.CommonTimestampsField
}

func (image *Image) Create() {
    database.DB.Create(&image)
}

func (image *Image) Save() (rowsAffected int64) {
    result := database.DB.Save(&image)
    return result.RowsAffected
}

func (image *Image) Delete() (rowsAffected int64) {
    result := database.DB.Delete(&image)
    return result.RowsAffected
}