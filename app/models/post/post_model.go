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
    AuthorID      string  `json:"author_id"`
    ColumnID      string  `json:"column_id,omitempty"`

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

func (postModel *Post) Updates(id,image,title,content string) (rowsAffected int64) {

    if image != "" && title != "" && content != ""{
        result := database.DB.Model(&postModel).
            Select("image_id","title","content").
            Updates(map[string]interface{}{"image_id":image,"title":title,"content":content})
        return result.RowsAffected
    }else if image != ""{
        result := database.DB.Model(&postModel).
            Select("image_id").
            Updates(map[string]interface{}{"image_id":image})
        return result.RowsAffected
    }else if title != ""{
        result := database.DB.Model(&postModel).
            Select("title").
            Updates(map[string]interface{}{"title":title})
        return result.RowsAffected
    }else if content != "" {
        result := database.DB.Model(&postModel).
            Select("content").
            Updates(map[string]interface{}{"content":content})
        return result.RowsAffected
    }
    return 0
}