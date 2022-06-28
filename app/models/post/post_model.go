package post

import (
    "GDColumn/app/models"
    "GDColumn/app/models/user"
    "GDColumn/pkg/database"
)

type Post struct {
    ID            string  `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
    Title         string  `json:"title,omitempty" `
    Content       string  `json:"content,omitempty" `
    Excerpt       string  `json:"excerpt,omitempty" `
    ImageID       string  `json:"-"`
    AuthorID      string  `json:"-"`
    ColumnID      string  `json:"columnId,omitempty"`

    Author        *user.User `json:"author,omitempty"`
    Image         *Image     `json:"image,omitempty"`

    models.CommonTimestampsField
}

type Image struct {
    ID  string `json:"id,omitempty"`
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

func (postModel *Post) Updates(image,title,content string) (rowsAffected int64) {
    var r = []rune(content)
    result := database.DB.Model(&postModel).
        First(&Post{}).Updates(Post{ImageID:image,Title:title,Content:content,Excerpt:string(r[:20])})
    return result.RowsAffected
}