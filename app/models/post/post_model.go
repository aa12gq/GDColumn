package post

import (
    "GDColumn/app/models"
    "GDColumn/app/models/user"
    "GDColumn/pkg/database"
)

type Post struct {
    ID            uint64  `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
    Title         string  `json:"title,omitempty" `
    Content       string  `json:"content,omitempty" `
    Excerpt       string  `json:"excerpt,omitempty" `
    ImageID       uint64  `json:"-"`
    AuthorID      uint64  `json:"-"`
    ColumnID      uint64  `json:"column_id,omitempty"`

    Author        *user.User `json:"author,omitempty"`
    Image         *Image     `json:"image,omitempty"`

    models.CommonTimestampsField
}

type Image struct {
    ID  uint64 `json:"id,omitempty"`
    URL string `json:"url,omitempty"`
}

func (post *Post) Create() {
    database.DB.Create(&post)
}

func (post *Post) Save() (rowsAffected int64) {
    result := database.DB.Save(&post)
    return result.RowsAffected
}

func (post *Post) Delete() (rowsAffected int64) {
    result := database.DB.Delete(&post)
    return result.RowsAffected
}