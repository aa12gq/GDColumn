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

func (columnModel *Column) Updates(id,image,title,des string) (rowsAffected int64) {

    if image != "" && title != "" && des != ""{
        result := database.DB.Model(&columnModel).
            Select("avatar_id","title","description").
            Updates(map[string]interface{}{"avatar_id":image,"title":title,"description":des})
        return result.RowsAffected
    }else if image != ""{
        result := database.DB.Model(&columnModel).
            Select("avatar_id").
            Updates(map[string]interface{}{"avatar_id":image})
        return result.RowsAffected
    }else if title != ""{
        result := database.DB.Model(&columnModel).
            Select("title").
            Updates(map[string]interface{}{"title":title})
        return result.RowsAffected
    }else if des != "" {
        result := database.DB.Model(&columnModel).
            Select("description").
            Updates(map[string]interface{}{"description":des})
        return result.RowsAffected
    }
    return 0
}