package post

import (
    "GDColumn/app/models"
    "GDColumn/app/models/column"
    "GDColumn/app/models/user"
    "GDColumn/pkg/database"
)

type Post struct {
    models.BaseModel

    Title         string  `json:"title,omitempty" `
    Content       string  `json:"content,omitempty" `
    Excerpt       string  `json:"excerpt,omitempty" `
    Image         string  `json:"image"`
    UserID        string  `json:"user_id,omitempty"`
    ColumnID      string  `json:"column_id,omitempty"`

    User          user.User `json:"user"`
    Column        column.Column  `json:"column"`

    models.CommonTimestampsField
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