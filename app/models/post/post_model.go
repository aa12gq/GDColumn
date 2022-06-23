package post

import (
    "GDColumn/app/models"
    "GDColumn/app/models/user"
    "GDColumn/pkg/database"
)

type Post struct {
    ID            uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
    Title         string  `json:"title,omitempty" `
    Content       string  `json:"content,omitempty" `
    Excerpt       string  `json:"excerpt,omitempty" `
    ImageID       string  `json:"image_id"`
    UserID        string  `json:"user_id,omitempty"`
    ColumnID      string  `json:"column_id,omitempty"`

    models.CommonTimestampsField
}

type Post2 struct {
    ID            uint64  `json:"id,omitempty"`
    Title         string  `json:"title,omitempty" `
    Content       string  `json:"content,omitempty" `
    Excerpt       string  `json:"excerpt,omitempty" `
    ImageID       string  `json:"image_id"`
    UserID        string  `json:"user_id,omitempty"`
    ColumnID      string  `json:"column_id,omitempty"`

    User          user.User
    Image         Image
    models.CommonTimestampsField
}

type Image struct {
    ID  uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
    URL string `json:"url"`
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